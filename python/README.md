# Python Notes

## Setup

- [Intro to pip and virtualenv](https://www.dabapps.com/blog/introduction-to-pip-and-virtualenv-python/)

This approach stores the complete Python environment, and all packages, into an 
arbitrary directory, eg:

```shell
$ pip install virtualenv
$ virtualenv venv
$ source venv/bin/activate
```

To deactivate a virtual env:

```shell
$ deactivate
```

To use a virtual env with Jupyter:
- Activate the virtual env, as shown above
- Run the following:

```shell
# $ pip install --user ipykernel
# Note: Above pip install failed when run from a virtual env with msg:
# ERROR: Can not perform a '--user' install. User site-packages are not visible in this virtualenv. 
# The version below seemed to work ok:
$ python -m ipykernel install --user --name=display_name 
Installed kernelspec myenv in /home/[user]/.local/share/jupyter/kernels/display_name
```

Note that the `--name=display_name` _is_ just a display name and the virtual env 
that is added is the currently activated one.
 
To list virtual envs available to jupyter:

```shell
$ jupyter kernelspec list
Available kernels:
  python3    /home/mike/.local/share/jupyter/kernels/python3
  venv       /home/mike/.local/share/jupyter/kernels/venv
```

To remove a virtual env from jupyter:

```shell
$ jupyter kernelspec uninstall display_name
```

Refs: 

- https://janakiev.com/blog/jupyter-virtual-envs/
- https://towardsdatascience.com/create-virtual-environment-using-virtualenv-and-add-it-to-jupyter-notebook-6e1bf4e03415


## Language

- [Closures](./closures/)
- [Context](./context/)
- [Dates](./dates/)
- [Decorators](./decorators/)
- [Docstrings](./docstrings/)
- [Exceptions](./exceptions/)
- [Files](./files/)
- [Generators](./generators/)
- [Idioms](./idioms/)
- [Iterators](./iterators/)
- [Lambdas](./lambdas/)
- [Object-Oriented Programming](./oop/)
- [Performance and efficiency](./performance/)
- [Strings and regular expressions](./strings/)


## Data Science

- [NumPy](./numpy/)
- [Pandas](./pandas/)
- See also [Data Science](../datasci/)

## Other resources 
- [DataQuest](https://dataquest.io) - data science course
- [Kaggle](https://www.kaggle.com) - hosted Jupyter notebooks, learning resources on machine learning, competitions
- [Colab](https://colab.research.google.com/) - browser-based python notebook, also links to lots of other courses
- [Jake VanderPlas](https://github.com/jakevdp) - Author of Python Data Science Handbook which is available on this repo

[Python Data Science Handbook](https://colab.research.google.com/github/jakevdp/PythonDataScienceHandbook/blob/master/notebooks/00.00-Preface.ipynb) - Open on Colab
