"""
TrainLoop Python SDK

This file ensures that 'trainloop' is treated as a Python package and
exports the Trainloop class for easy import.
"""

from .client import Trainloop, SampleFeedbackType

__all__ = ["Trainloop", "SampleFeedbackType"]
