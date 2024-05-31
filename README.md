# DevicesTransmissionSimulation

Simulating a mobile network transmission scenario where \( n \) users try to transmit using one of \( M \) preambles of a base station.

This repository contains a simulation of a mobile network transmission scenario. The simulation models how multiple user devices (UEs) attempt to access the network using a limited number of preambles. By simulating this process, we can analyze the performance, collision rates, and other metrics crucial for optimizing network access procedures.

## Requirements

* Golang >= 1.22
* Python >= 3.8

## Setup

After installing the requirements, Jupyter Notebook must be installed to run the notebook with the models comparison.

```
pip install -r requirements.txt
```

## Running

The transmission simulation is written in Golang. To run the monte carlo simulator:

```
go run cmd/monte_carlo/main.go -r 6 -n 30 -M 64
```

The code must be executed to generate the simulation results. After that, run Jupyter Notebook to visualize the comparison notebook:

```
jupyter notebook
```

Then, select the file `MonteCarloXAnalytical.ipynb` in the notebooks directory to visualize the results of the simulation.

### Examples

* Running 10⁶ experiments with 30 devices and 10 preambles:
```
go run cmd/monte_carlo/main.go -r 6 -n 30 -M 10
```