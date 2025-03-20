"""
TrainLoop Python SDK

This file ensures that 'trainloop' is treated as a Python package and
exports the Client class for easy import.
"""

from .trainloop import Client, SampleFeedbackType

__all__ = ["Client", "SampleFeedbackType"]
