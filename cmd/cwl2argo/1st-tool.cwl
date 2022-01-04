cwlVersion: v1.2
class: CommandLineTool
requirements:
  - class: DockerRequirement 
    dockerPull: ubuntu:20.04
baseCommand: echo
id: echo-tool
inputs:
  message:
    type: string
    inputBinding:
      position: 1
outputs: []
