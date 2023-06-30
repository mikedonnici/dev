# Github actions

- workflow automation
  - continuous integration - continuous delivery
  - code and repository management

## Key components

- *repository*
  - *workflows* 
    - triggered by events
    - contain one or more jobs
      - *jobs* 
        - define a runner (execution environment) - _every job has its own runner_
        - can be conditional
        - run in parallel by default, can be run sequentially
        - contain one or more steps
        - *steps* 
          - can be conditional
          - run sequentially 
          - execute a shell script (command) or action
          - *actions*
            - own or third party  

## `run` commands

- A step can run a shell command using the `run` keyword
- Multiple or multiline commands use the pipe symbol

```yaml
steps: 
  name: Step one 
  run: |
    echo "one" 
    echo "two"
    echo "three"
```

## `uses` action `with` configuration

- a step can run an action with the `uses` keyword
- configuration (key-val pairs) for the action is provided using the `with` keyword 
- *note: configuration keys must be supported by the action or it will fail*

```yaml
steps: 
  name: Step two
  uses: some-action@v1 
  with: 
    key1: val1 
    key2: val2
```

## Expressions and context objects 

- Enables access to _metadata_ such as env vars
- Available metadata is collectively referred to as *context*
- Access context vars using the _expression_ syntax: `${{ }}`
- *Expressions* are combinations of _literals_, context references and _functions_
- ref: 
  - https://docs.github.com/en/actions/learn-github-actions/contexts
  - https://docs.github.com/en/actions/learn-github-actions/expressions


```yaml
name: output values
on: workflow_dispatch
jobs:
  output:
    runs-on: ubuntu-latest
    steps:
      - name: Output github key context values
        run: echo "${{ toJSON(github) }}"
```



