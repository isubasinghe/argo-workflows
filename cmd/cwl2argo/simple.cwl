cwlVersion: v1.0
class: CommandLineTool
baseCommand: 
  - tar
  - --extract
  - -f
id: echo-tool

requirements:
  - class: DockerRequirement 
    dockerPull: ubuntu:20.04
    dockerOutputDirectory: /mnt/pvol 

  - class: ResourceRequirement 
    outdirMin: 1Gi

inputs:
  tarfile:
    type: File
    inputBinding:
      position: 0

outputs: 
  extracted_file:
    type: File 
    outputBinding:
      glob: hello.txt
