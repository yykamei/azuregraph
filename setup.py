# -*- coding: utf-8 -*-
from setuptools import setup, find_packages

try:
    long_description = open("README.rst").read()
except IOError:
    long_description = ""

setup(
    name="azuregraph",
    version="0.1.0",
    description="Azure AD Graph API tool",
    license="MIT",
    author="Yutaka Kamei",
    packages=find_packages(),
    scripts=[
        'azuregraph-tool'
    ],
    long_description=long_description,
    classifiers=[
        "Programming Language :: Python",
        "Programming Language :: Python :: 3",
    ]
)
