cwlVersion: v1.0
class: CommandLineTool
baseCommand: 
  - echo
id: echo-tool
requirements:
  - class: DockerRequirement 
    dockerPull: ubuntu:20.04

inputs:
  example_file:
    type: File
    inputBinding:
      position: 0
outputs: []
