module.exports = {
  docs: [
    "introduction",
    {
      type: "category",
      label: "Setup",
      items: [
        "setup/bare-metal",
        "setup/azure-pipelines",
        "setup/docker",
        "setup/gitlab",
        "setup/kubernetes",
        "setup/travis"
      ]
    },
    {
      type: "category",
      label: "Training",
      items: [
        "training/hello-world",
        "training/inputs-outputs",
        {
          type: "category",
          label: "Portable Dev",
          items: [
            "run-a-go-service",
            "run-a-react-app"
          ]
        },
        {
          type: "category",
          label: "Containers",
          items: [
            "training/containers/how-do-i-communicate-with-an-opctl-container",
            "training/containers/how-do-i-get-opctl-containers-to-communicate"
          ]
        },
        {
          type: "category",
          label: "UI",
          items: [
            "training/ui/how-do-i-visualize-an-op"
          ]
        }
      ]
    },
    {
      type: "category",
      label: "Reference",
      items: [
        {
          type: "category",
          label: "Opspec",
          items: [
            "reference/opspec/index",
            {
              type: "category",
              label: "Op [directory]",
              items: [
                "reference/opspec/op-directory/index",
                {
                  type: "category",
                  label: "Op [object]",
                  items: [
                    "reference/opspec/op-directory/op/index",
                    {
                      type: "category",
                      label: "Call [object]",
                      items: [
                        "reference/opspec/op-directory/op/call/index",
                        {
                          type: "category",
                          label: "Container Call [object]",
                          items: [
                            "reference/opspec/op-directory/op/call/container/index",
                            "reference/opspec/op-directory/op/call/container/image",
                          ]
                        },
                        "reference/opspec/op-directory/op/call/loop-vars",
                        "reference/opspec/op-directory/op/call/op",
                        "reference/opspec/op-directory/op/call/parallel-loop",
                        "reference/opspec/op-directory/op/call/predicate",
                        "reference/opspec/op-directory/op/call/pull-creds",
                        "reference/opspec/op-directory/op/call/rangeable-value",
                        "reference/opspec/op-directory/op/call/serial-loop"
                      ]
                    },
                    {
                      type: "category",
                      label: "Parameter [object]",
                      items: [
                        "reference/opspec/op-directory/op/parameter/index",
                        "reference/opspec/op-directory/op/parameter/array",
                        "reference/opspec/op-directory/op/parameter/boolean",
                        "reference/opspec/op-directory/op/parameter/dir",
                        "reference/opspec/op-directory/op/parameter/file",
                        "reference/opspec/op-directory/op/parameter/number",
                        "reference/opspec/op-directory/op/parameter/object",
                        "reference/opspec/op-directory/op/parameter/socket",
                        "reference/opspec/op-directory/op/parameter/string",
                      ]
                    },
                    "reference/opspec/op-directory/op/identifier",
                    "reference/opspec/op-directory/op/initializer",
                    "reference/opspec/op-directory/op/markdown",
                    "reference/opspec/op-directory/op/variable-reference"
                  ]
                }
              ]
            },
            {
              type: "category",
              label: "Types",
              items: [
                "reference/opspec/types/array",
                "reference/opspec/types/boolean",
                "reference/opspec/types/dir",
                "reference/opspec/types/file",
                "reference/opspec/types/number",
                "reference/opspec/types/object",
                "reference/opspec/types/socket",
                "reference/opspec/types/string",
              ]
            }
          ]
        },
        'reference/ui',
        {
          type: "category",
          label: "CLI",
          items: [
            "reference/cli/global-options",
            "reference/cli/auth",
            "reference/cli/events",
            "reference/cli/ls",
            "reference/cli/node",
            "reference/cli/op",
            "reference/cli/run",
            "reference/cli/self-update",
            "reference/cli/ui",
          ]
        },
        {
          type: "category",
          label: "SDKs",
          items: [
            {
              type: "link",
              label: "Go",
              href: "https://github.com/opctl/opctl/tree/master/sdks/go"
            },
            {
              type: "link",
              label: "Js",
              href: "https://github.com/opctl/opctl/tree/master/sdks/js"
            },
            {
              type: "link",
              label: "React",
              href: "https://github.com/opctl/opctl/tree/master/sdks/react"
            }
          ]
        }
      ]
    }
  ]
};
