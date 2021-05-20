# Environment Variables

## Objective 
I would like to take a single config-like file and load it for each environment. 

## Variable Sources 
1. `shared.yaml`: variables that all environments share, but may want to be changed on a per-project basis.
2. `dev.yaml`: variables for local development environment. 
3. `prod.yaml`: variables for external production environment. 