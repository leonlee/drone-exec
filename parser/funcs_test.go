package parser

import (
	"testing"

	"github.com/franela/goblin"
)

func Test_Funcs(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("ImageName", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
			})

			g.Describe("that is a NodeCompose", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeCompose
				})

				g.Describe("without an Image", func() {
					g.It("returns image missing error", func() {
						g.Assert(ImageName(node)).Equal(ErrImageMissing)
					})
				})

				g.Describe("with an Image", func() {
					g.Describe("that has a tag", func() {
						g.BeforeEach(func() {
							node.Image = `fakeImage:latest`
						})

						g.It("does not alter the node's Image", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`fakeImage:latest`)
						})
					})

					g.Describe("that does not have a tag", func() {
						g.BeforeEach(func() {
							node.Image = `fakeImage`
						})

						g.It("uses the latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`fakeImage:latest`)
						})
					})
				})
			})

			g.Describe("that is a NodeBuild", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeBuild
				})

				g.Describe("without an Image", func() {
					g.Describe("without Commands", func() {
						g.It("does not alter the node's Image", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(``)
						})

						g.It("does not return an error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("with Commands", func() {
						g.BeforeEach(func() {
							node.Commands = []string{`do stuff`}
						})

						g.It("does not alter the node's Image", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(``)
						})

						g.It("returns image missing error", func() {
							g.Assert(ImageName(node)).Equal(ErrImageMissing)
						})
					})
				})

				g.Describe("with an Image", func() {
					g.Describe("that has a tag", func() {
						g.BeforeEach(func() {
							node.Image = `fakeImage:latest`
						})

						g.It("does not alter the node's Image", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`fakeImage:latest`)
						})
					})

					g.Describe("that does not have a tag", func() {
						g.BeforeEach(func() {
							node.Image = `fakeImage`
						})

						g.It("uses the latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`fakeImage:latest`)
						})
					})
				})
			})

			g.Describe("that is a NodeClone", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeClone
				})

				g.Describe("without an Image", func() {
					g.It("sets the Image to DefaultCloner", func() {
						ImageName(node)
						g.Assert(node.Image).Equal(`plugins/drone-git:latest`)
					})

					g.It("does not error", func() {
						g.Assert(ImageName(node)).Equal(nil)
					})
				})

				g.Describe("with an Image", func() {
					g.Describe("that has a plugin path", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-mycloner`
						})

						g.It("adds the latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycloner:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that is a basic plugin name", func() {
						g.BeforeEach(func() {
							node.Image = `mycloner`
						})

						g.It("adds the plugin prefix and latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycloner:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that uses underscores", func() {
						g.BeforeEach(func() {
							node.Image = `my_cloner`
						})

						g.It("converts the underscore to a dash", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-my-cloner:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})
				})
			})

			g.Describe("that is a NodeCache", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeCache
				})

				g.Describe("without an Image", func() {
					g.It("sets the Image to NodeCache", func() {
						ImageName(node)
						g.Assert(node.Image).Equal(`plugins/drone-cache:latest`)
					})

					g.It("does not error", func() {
						g.Assert(ImageName(node)).Equal(nil)
					})
				})

				g.Describe("with an Image", func() {
					g.Describe("that has a plugin path", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-mycacher`
						})

						g.It("adds the latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that is a basic plugin name", func() {
						g.BeforeEach(func() {
							node.Image = `mycacher`
						})

						g.It("adds the plugin prefix and latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that uses underscores", func() {
						g.BeforeEach(func() {
							node.Image = `my_cacher`
						})

						g.It("converts the underscore to a dash", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-my-cacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})
				})
			})

			g.Describe("that is some other node", func() {
				g.BeforeEach(func() {
					node.NodeType = fakeNodeType
				})

				g.Describe("without an Image", func() {
					g.It("does not error", func() {
						// I suspect this should actually error rather than
						// setting the image to `plugins/drone-:latest`
						g.Assert(ImageName(node)).Equal(nil)
					})
				})

				g.Describe("with an Image", func() {
					g.Describe("that has a plugin path", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-mycacher`
						})

						g.It("adds the latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that is a basic plugin name", func() {
						g.BeforeEach(func() {
							node.Image = `mycacher`
						})

						g.It("adds the plugin prefix and latest tag", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-mycacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})

					g.Describe("that uses underscores", func() {
						g.BeforeEach(func() {
							node.Image = `my_cacher`
						})

						g.It("converts the underscore to a dash", func() {
							ImageName(node)
							g.Assert(node.Image).Equal(`plugins/drone-my-cacher:latest`)
						})

						g.It("does not error", func() {
							g.Assert(ImageName(node)).Equal(nil)
						})
					})
				})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(ImageName(node)).Equal(nil)
			})
		})
	})

	g.Describe("ImageMatch", func() {
		// TODO
	})

	g.Describe("ImagePull", func() {
		var pull bool

		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
			})

			g.Describe("that is a NodeBuild", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeBuild
					pull = true
				})

				g.It("doesn't set Pull", func() {
					ImagePull(node, pull)
					g.Assert(node.Pull).Equal(false)
				})
			})

			g.Describe("that is a NodeCompose", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeCompose
					pull = true
				})

				g.It("doesn't set Pull", func() {
					ImagePull(node, pull)
					g.Assert(node.Pull).Equal(false)
				})
			})

			g.Describe("that is some other node", func() {
				g.BeforeEach(func() {
					node.NodeType = fakeNodeType
				})

				g.Describe("when pull is true", func() {
					g.BeforeEach(func() {
						pull = true
					})

					g.It("sets Pull to true", func() {
						ImagePull(node, pull)
						g.Assert(node.Pull).Equal(true)
					})

					g.It("does not error", func() {
						g.Assert(ImagePull(node, pull)).Equal(nil)
					})
				})

				g.Describe("when pull is false", func() {
					g.BeforeEach(func() {
						pull = false
					})

					g.It("sets Pull to false", func() {
						ImagePull(node, pull)
						g.Assert(node.Pull).Equal(false)
					})

					g.It("does not error", func() {
						g.Assert(ImagePull(node, pull)).Equal(nil)
					})
				})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(ImagePull(node, pull)).Equal(nil)
			})
		})
	})

	g.Describe("Sanitize", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
				node.Privileged = true
				node.Volumes = []string{`some stuff`}
				node.Devices = []string{`some devices`}
				node.Net = `someNet`
				node.Entrypoint = []string{`super entrypoint`}
			})

			g.It("sanitizes the node", func() {
				Sanitize(node)
				g.Assert(node.Privileged).Equal(false)
				g.Assert(node.Volumes).Equal([]string(nil))
				g.Assert(node.Devices).Equal([]string(nil))
				g.Assert(node.Net).Equal(``)
				g.Assert(node.Entrypoint).Equal([]string{})
			})

			g.It("does not error", func() {
				g.Assert(Sanitize(node)).Equal(nil)
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(Sanitize(node)).Equal(nil)
			})
		})
	})

	g.Describe("Escalate", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
				node.Privileged = false
				node.Volumes = []string{`some stuff`}
				node.Net = `someNet`
				node.Entrypoint = []string{`super entrypoint`}
			})

			g.Describe("that is a NodePublish", func() {
				g.BeforeEach(func() {
					node.NodeType = NodePublish
				})

				g.Describe("and whitelisted", func() {
					g.Describe("plugins/drone-docker", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-docker`
						})

						g.It("escalates the node", func() {
							Escalate(node)
							g.Assert(node.Privileged).Equal(true)
							g.Assert(node.Volumes).Equal([]string(nil))
							g.Assert(node.Net).Equal(``)
							g.Assert(node.Entrypoint).Equal([]string{})
						})

						g.It("does not error", func() {
							g.Assert(Escalate(node)).Equal(nil)
						})
					})

					g.Describe("plugins/drone-gcr", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-gcr`
						})

						g.It("escalates the node", func() {
							Escalate(node)
							g.Assert(node.Privileged).Equal(true)
							g.Assert(node.Volumes).Equal([]string(nil))
							g.Assert(node.Net).Equal(``)
							g.Assert(node.Entrypoint).Equal([]string{})
						})

						g.It("does not error", func() {
							g.Assert(Escalate(node)).Equal(nil)
						})
					})

					g.Describe("plugins/drone-ecr", func() {
						g.BeforeEach(func() {
							node.Image = `plugins/drone-ecr`
						})

						g.It("escalates the node", func() {
							Escalate(node)
							g.Assert(node.Privileged).Equal(true)
							g.Assert(node.Volumes).Equal([]string(nil))
							g.Assert(node.Net).Equal(``)
							g.Assert(node.Entrypoint).Equal([]string{})
						})

						g.It("does not error", func() {
							g.Assert(Escalate(node)).Equal(nil)
						})
					})
				})

				g.Describe("and not whitelisted", func() {
					g.BeforeEach(func() {
						node.NodeType = NodePublish
						node.Image = ``
					})

					g.It("does not alter the node", func() {
						Escalate(node)
						g.Assert(node.Privileged).Equal(false)
						g.Assert(node.Volumes).Equal([]string{`some stuff`})
						g.Assert(node.Net).Equal(`someNet`)
						g.Assert(node.Entrypoint).Equal([]string{`super entrypoint`})
					})

					g.It("does not error", func() {
						g.Assert(Escalate(node)).Equal(nil)
					})
				})
			})

			g.Describe("that is not NodePublish", func() {
				g.BeforeEach(func() {
					node.NodeType = fakeNodeType
				})

				g.It("does not error", func() {
					g.Assert(Escalate(node)).Equal(nil)
				})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(Escalate(node)).Equal(nil)
			})
		})
	})

	g.Describe("DefaultNotifyFilter", func() {
		// TODO
	})

	g.Describe("HttpProxy", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
			})

			g.It("does not error", func() {
				g.Assert(HttpProxy(node)).Equal(nil)
			})

			// Note: should test that proxy related ENV vars are added to node but
			// difficult to stub os.Environ()
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(HttpProxy(node)).Equal(nil)
			})
		})
	})

	g.Describe("Cache", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
			})

			g.Describe("that is a NodeCache", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeCache
				})

				g.It("adds the /cache volume", func() {
					Cache(node, `somedir`)
					g.Assert(node.Volumes).Eql([]string{`/var/lib/drone/cache/somedir:/cache`})
				})

				g.It("does not error", func() {
					g.Assert(Cache(node, `somedir`)).Equal(nil)
				})
			})

			g.Describe("that is some other node", func() {
				g.BeforeEach(func() {
					node.NodeType = fakeNodeType
				})

				g.It("does not alter the volumes", func() {
					Cache(node, `somedir`)
					g.Assert(node.Volumes).Eql([]string(nil))
				})

				g.It("does not error", func() {
					g.Assert(Cache(node, `somedir`)).Equal(nil)
				})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(Cache(node, `somedir`)).Equal(nil)
			})
		})
	})

	g.Describe("Debug", func() {
		g.Describe("with a DockerNode", func() {
			var node *DockerNode

			g.BeforeEach(func() {
				node = &DockerNode{}
			})

			g.Describe("that is a NodeCache", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeCache
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is a NodeClone", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeClone
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is a NodeDeploy", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeDeploy
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is a NodeNotify", func() {
				g.BeforeEach(func() {
					node.NodeType = NodeNotify
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is a NodePublish", func() {
				g.BeforeEach(func() {
					node.NodeType = NodePublish
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is a NodePublish", func() {
				g.BeforeEach(func() {
					node.NodeType = NodePublish
				})

				g.It("adds debug flag to the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				// Note: the `debug` parameter is ignored
				g.It("adds debug flag to the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string{`DEBUG=true`})
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})

			g.Describe("that is some other node", func() {
				g.BeforeEach(func() {
					node.NodeType = fakeNodeType
				})

				g.It("does not alter the environment", func() {
					Debug(node, true)
					g.Assert(node.Environment).Eql([]string(nil))
				})

				// Note: the `debug` parameter is ignored
				g.It("does not alter the environment", func() {
					Debug(node, false)
					g.Assert(node.Environment).Eql([]string(nil))
				})

				g.It("does not error", func() {
					g.Assert(Debug(node, true)).Equal(nil)
				})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(Debug(node, true)).Equal(nil)
			})
		})
	})

	g.Describe("Mount", func() {
		var from, to string

		g.Describe("with a DockerNode", func() {
			var node *DockerNode
			g.BeforeEach(func() {
				node = &DockerNode{}
				from = `source/dir`
				to = `dest/dir`
			})

			g.It("does not error", func() {
				g.Assert(Mount(node, from, to)).Equal(nil)
			})

			g.It("adds the volume mount", func() {
				Mount(node, from, to)
				g.Assert(node.Volumes).Equal([]string{`source/dir:dest/dir`})
			})
		})

		g.Describe("with a different Node", func() {
			var node Node

			g.BeforeEach(func() {
				node = &fakeNode{}
			})

			// Note: this should probably return a type error
			g.It("does not error", func() {
				g.Assert(Mount(node, from, to)).Equal(nil)
			})
		})
	})
}

const fakeNodeType = 100

type fakeNode struct{}

func (n *fakeNode) Type() NodeType {
	return fakeNodeType
}
