# Kinney

The electric grid, composed of substations, transmission lines, transformers and
more, delivers electricity from the power plants to our homes and businesses.
Its transformation to a "Smart Grid" requires real time sensing of energy
production and consumption, anticipating usage needs, and smoothing load
fluctuation by way of curtailing demand and bringing cleaner Distributed Energy
Resources (DERs) online.  This transformation is a fascinating challenge due to
the need to integrate legacy system, varied protocols, and myraid devices, while
securing remote dispersed systems, protecting user data, coping with
unpredictable DERs, and incentivizing changes in human behavior: all to deliver
reliable, efficient, and progressively cleaner power to meet our needs.

This project defines models for devices that produce and consume energy,
incorporating environmental factors that influence consumption such as weather,
time of day, and seasonal fluctuations, and providing mechanisms for real time
monitoring and control.

Together, through open standards and implementations, we can achieve this Smart
Grid, reduce global warming and help preserve our beautiful planet.


## Development


### Python

1. Python development on this project requires [`pipenv`] and [`pyenv`].  For
   example, on MacOS with Homebrew, these can be installed with:

   ```bash
   $ brew install pipenv pyenv
   ```

1. Install the pinned version of Python with `pyenv`.  The current version can
   be found in the [`Pipfile`](Pipfile) under `requires.python_full_version`.

   ```bash
   $ pyenv install <version>
   ```

1. Set up the Python virtual environment for local development:

   ```bash
   $ make pipenv
   ```

1. `pipenv` can now be used to run the Python code.  For example:

   ```bash
   $ pipenv run python -m kinney.orchestrator.server
   ```

   To get an interactive Python shell, use `pipenv run python`.

[`pipenv`]: https://github.com/pypa/pipenv
[`pyenv`]: https://github.com/pyenv/pyenv


### Protocol Buffers

The generated Protocol Buffer sources can be refreshed with:

```bash
$ make protos
```

The generated code is checked in so that only developers who need to change the
API definitions will need to have the Protocol Buffer compiled installed.
