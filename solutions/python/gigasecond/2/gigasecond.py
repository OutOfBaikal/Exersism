"""Gigasecond"""
from datetime import timedelta

def add(moment):
    """add 1.000.000.000 seconds to moment"""
    return moment + timedelta(seconds=1000000000)
