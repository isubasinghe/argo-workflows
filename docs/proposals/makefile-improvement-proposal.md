# Proposal for Makefile improvements. 

## Introduction 
The motivation for this proposal is to enable developers working on Argo Workflows to use build tools in a more reproducable way. 
Currently the Makefile is unfortunately too opinionated and as a result is often a blocker when first setting up Argo Workflows locally. 
I believe we should shrink the responsibilities of the Makefile and where possible outsource areas of responsibility to more specialised technology, such 
as Devenv/Nix in the case of dependency management. 

## Proposal Specifics
In order to better address reproducability, it is better to split up the duties the Makefile currently performs into various sub components, that can be assembled in more appropriate technology. 

### Current Makefile Responsibilities


### Proposed Makefile Responsibilities


### Devenv
