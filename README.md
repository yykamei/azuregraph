# Azure AD Graph API Tool

## Installation

    $ git clone https://github.com/yykamei/azuregraph /path/to/azuregraph

## Preparation

Firstly, you must set up to use Azure AD Graph API.
See Step 3 of https://github.com/Azure-Samples/active-directory-dotnet-graphapi-console .

## Usage

Create a file like this, which includes Azure AD Graph API connection informations.
This file is named *azure.info* .

    [azure]
    tenant_id = example.com
    client_id = xxx
    client_secret = xxx

NOTE: You should set appropriate file permission because this file has credentials.

Then, you can dispatch Azure Graph API request.
Following command means "Get users".

    $ /path/to/azuregraph/azuregraph-tool -I azure.info list users
