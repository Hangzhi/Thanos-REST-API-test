# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesGetRules**](DefaultApi.md#RulesGetRules) | **Get** /api/v1/rules | Gets rules. User can pass \&quot;rules?type&#x3D;record\&quot; to get rule record instead of full rules.
[**RulesGetRulesAlert**](DefaultApi.md#RulesGetRulesAlert) | **Get** /api/v1/alerts | Gets alerts. No parameters will be processed.

# **RulesGetRules**
> RuleGroups RulesGetRules(ctx, optional)
Gets rules. User can pass \"rules?type=record\" to get rule record instead of full rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRulesGetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRulesGetRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| / The type to request. | 

### Return type

[**RuleGroups**](RuleGroups.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RulesGetRulesAlert**
> Alert RulesGetRulesAlert(ctx, )
Gets alerts. No parameters will be processed.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Alert**](Alert.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

