# Azure AD Graph API Tool

azuregraph-tool dispatches [Azure AD Graph API][azure_ad_graph_api].
This supports api-version *1.6*, and only GET APIs are supported.

[azure_ad_graph_api]: https://azure.microsoft.com/en-us/documentation/articles/active-directory-graph-api/

## Installation

    $ git clone https://github.com/yykamei/azuregraph-tool /path/to/azuregraph-tool

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

    $ /path/to/azuregraph-tool/azuregraph-tool -I azure.info list users
