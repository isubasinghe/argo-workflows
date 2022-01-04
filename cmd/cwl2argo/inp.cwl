class: Workflow
doc: "Reverse the lines in a document, then sort those lines."
cwlVersion: v1.0

hints:
  - class: DockerRequirement
    dockerPull: docker.io/debian:8

inputs:
  workflow_input:
    type: File
    doc: "The input file to be processed."
    format: iana:text/plain
    default:
      class: File
      location: hello.txt
  reverse_sort:
    type: boolean
    default: true
    doc: "If true, reverse (decending) sort"

outputs:
  sorted_output:
    type: File
    outputSource: sorted/sorted_output
    doc: "The output with the lines reversed and sorted."

steps:
  rev:
    in:
      revtool_input: workflow_input
    out: [revtool_output]
    run: revtool.cwl

  sorted:
    in:
      sorted_input: rev/revtool_output
      reverse: reverse_sort
    out: [sorted_output]
    run: sorttool.cwl
