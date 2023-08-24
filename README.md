# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.6
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "//"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://:443*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AudioAPI* | [**ApiAudioEqualizerGet**](docs/AudioAPI.md#apiaudioequalizerget) | **Get** /api/audio/equalizer | Get the current equalizer settings
*AudioAPI* | [**ApiAudioEqualizerPut**](docs/AudioAPI.md#apiaudioequalizerput) | **Put** /api/audio/equalizer | Set the current equalizer settings
*AudioAPI* | [**ApiAudioInputsDanteReferenceGet**](docs/AudioAPI.md#apiaudioinputsdantereferenceget) | **Get** /api/audio/inputs/dante/reference | Get settings of the reference input
*AudioAPI* | [**ApiAudioInputsDanteReferencePut**](docs/AudioAPI.md#apiaudioinputsdantereferenceput) | **Put** /api/audio/inputs/dante/reference | Set reference input settings
*AudioAPI* | [**ApiAudioInputsMicrophoneBeamDirectionGet**](docs/AudioAPI.md#apiaudioinputsmicrophonebeamdirectionget) | **Get** /api/audio/inputs/microphone/beam/direction | Get the current direction of the beam
*AudioAPI* | [**ApiAudioInputsMicrophoneBeamGet**](docs/AudioAPI.md#apiaudioinputsmicrophonebeamget) | **Get** /api/audio/inputs/microphone/beam | Get the current beam settings
*AudioAPI* | [**ApiAudioInputsMicrophoneBeamPut**](docs/AudioAPI.md#apiaudioinputsmicrophonebeamput) | **Put** /api/audio/inputs/microphone/beam | Set the current beam settings
*AudioAPI* | [**ApiAudioInputsMicrophoneExclusionZonesGet**](docs/AudioAPI.md#apiaudioinputsmicrophoneexclusionzonesget) | **Get** /api/audio/inputs/microphone/exclusionZones | Get the supported exclusion zone ids
*AudioAPI* | [**ApiAudioInputsMicrophoneExclusionZonesIdGet**](docs/AudioAPI.md#apiaudioinputsmicrophoneexclusionzonesidget) | **Get** /api/audio/inputs/microphone/exclusionZones/{id} | Get the current exclusion zone settings of zone number &#x60;id&#x60;
*AudioAPI* | [**ApiAudioInputsMicrophoneExclusionZonesIdPut**](docs/AudioAPI.md#apiaudioinputsmicrophoneexclusionzonesidput) | **Put** /api/audio/inputs/microphone/exclusionZones/{id} | Set the current exclusion zone settings of zone number &#x60;id&#x60;
*AudioAPI* | [**ApiAudioInputsMicrophoneLevelGet**](docs/AudioAPI.md#apiaudioinputsmicrophonelevelget) | **Get** /api/audio/inputs/microphone/level | Get the current microphone input level
*AudioAPI* | [**ApiAudioInputsMicrophonePriorityZonesGet**](docs/AudioAPI.md#apiaudioinputsmicrophonepriorityzonesget) | **Get** /api/audio/inputs/microphone/priorityZones | Get the supported priority zone ids
*AudioAPI* | [**ApiAudioInputsMicrophonePriorityZonesIdGet**](docs/AudioAPI.md#apiaudioinputsmicrophonepriorityzonesidget) | **Get** /api/audio/inputs/microphone/priorityZones/{id} | Get the current priority zone settings of zone number &#x60;id&#x60;
*AudioAPI* | [**ApiAudioInputsMicrophonePriorityZonesIdPut**](docs/AudioAPI.md#apiaudioinputsmicrophonepriorityzonesidput) | **Put** /api/audio/inputs/microphone/priorityZones/{id} | Set the current priority zone settings of zone number &#x60;id&#x60;
*AudioAPI* | [**ApiAudioInputsReferenceLevelGet**](docs/AudioAPI.md#apiaudioinputsreferencelevelget) | **Get** /api/audio/inputs/reference/level | Get the current level of the digital reference input
*AudioAPI* | [**ApiAudioNoiseGateGet**](docs/AudioAPI.md#apiaudionoisegateget) | **Get** /api/audio/noiseGate | Get the current noise gate settings
*AudioAPI* | [**ApiAudioNoiseGatePut**](docs/AudioAPI.md#apiaudionoisegateput) | **Put** /api/audio/noiseGate | Set the current noise gate settings
*AudioAPI* | [**ApiAudioOutputsAnalogGet**](docs/AudioAPI.md#apiaudiooutputsanalogget) | **Get** /api/audio/outputs/analog | Get the current settings of the analog output
*AudioAPI* | [**ApiAudioOutputsAnalogPut**](docs/AudioAPI.md#apiaudiooutputsanalogput) | **Put** /api/audio/outputs/analog | Set the analog output settings
*AudioAPI* | [**ApiAudioOutputsDanteFarEndGet**](docs/AudioAPI.md#apiaudiooutputsdantefarendget) | **Get** /api/audio/outputs/dante/farEnd | Get the current settings of the far end output
*AudioAPI* | [**ApiAudioOutputsDanteFarEndPut**](docs/AudioAPI.md#apiaudiooutputsdantefarendput) | **Put** /api/audio/outputs/dante/farEnd | Set the far end output settings
*AudioAPI* | [**ApiAudioOutputsDanteLocalGet**](docs/AudioAPI.md#apiaudiooutputsdantelocalget) | **Get** /api/audio/outputs/dante/local | Get the current settings of the local output
*AudioAPI* | [**ApiAudioOutputsDanteLocalPut**](docs/AudioAPI.md#apiaudiooutputsdantelocalput) | **Put** /api/audio/outputs/dante/local | Set the local output settings
*AudioAPI* | [**ApiAudioOutputsGlobalMuteGet**](docs/AudioAPI.md#apiaudiooutputsglobalmuteget) | **Get** /api/audio/outputs/global/mute | Get the mute status of the device
*AudioAPI* | [**ApiAudioOutputsGlobalMutePut**](docs/AudioAPI.md#apiaudiooutputsglobalmuteput) | **Put** /api/audio/outputs/global/mute | Mute all audio outputs
*AudioAPI* | [**ApiAudioRoomInUseActivityLevelGet**](docs/AudioAPI.md#apiaudioroominuseactivitylevelget) | **Get** /api/audio/roomInUse/activityLevel | Get the current room in use activity level
*AudioAPI* | [**ApiAudioRoomInUseConfigGet**](docs/AudioAPI.md#apiaudioroominuseconfigget) | **Get** /api/audio/roomInUse/config | Get the configuration of the room in use feature
*AudioAPI* | [**ApiAudioRoomInUseGet**](docs/AudioAPI.md#apiaudioroominuseget) | **Get** /api/audio/roomInUse | Get the current room in use state
*AudioAPI* | [**ApiAudioVoiceLiftGet**](docs/AudioAPI.md#apiaudiovoiceliftget) | **Get** /api/audio/voiceLift | Get the current voice lift settings
*AudioAPI* | [**ApiAudioVoiceLiftPut**](docs/AudioAPI.md#apiaudiovoiceliftput) | **Put** /api/audio/voiceLift | Set the current voice lift settings
*DeviceAPI* | [**ApiDeviceIdentificationGet**](docs/DeviceAPI.md#apideviceidentificationget) | **Get** /api/device/identification | Get the state of device identification
*DeviceAPI* | [**ApiDeviceIdentificationPut**](docs/DeviceAPI.md#apideviceidentificationput) | **Put** /api/device/identification | Set the state of device identification
*DeviceAPI* | [**ApiDeviceIdentityGet**](docs/DeviceAPI.md#apideviceidentityget) | **Get** /api/device/identity | Get the device identity
*DeviceAPI* | [**ApiDeviceLedsRingGet**](docs/DeviceAPI.md#apideviceledsringget) | **Get** /api/device/leds/ring | Get the current led ring settings
*DeviceAPI* | [**ApiDeviceLedsRingPut**](docs/DeviceAPI.md#apideviceledsringput) | **Put** /api/device/leds/ring | Set the current led ring brightness and colors
*DeviceAPI* | [**ApiDevicePowerPoeDaisychainGet**](docs/DeviceAPI.md#apidevicepowerpoedaisychainget) | **Get** /api/device/power/poe/daisychain | Get the daisy chain and PoE state
*DeviceAPI* | [**ApiDeviceSiteGet**](docs/DeviceAPI.md#apidevicesiteget) | **Get** /api/device/site | Get the device site information
*DeviceAPI* | [**ApiDeviceStateGet**](docs/DeviceAPI.md#apidevicestateget) | **Get** /api/device/state | Get the device state
*FastResourceAPI* | [**ApiAudioInputsMicrophoneBeamDirectionGet**](docs/FastResourceAPI.md#apiaudioinputsmicrophonebeamdirectionget) | **Get** /api/audio/inputs/microphone/beam/direction | Get the current direction of the beam
*FastResourceAPI* | [**ApiAudioInputsMicrophoneLevelGet**](docs/FastResourceAPI.md#apiaudioinputsmicrophonelevelget) | **Get** /api/audio/inputs/microphone/level | Get the current microphone input level
*FastResourceAPI* | [**ApiAudioInputsReferenceLevelGet**](docs/FastResourceAPI.md#apiaudioinputsreferencelevelget) | **Get** /api/audio/inputs/reference/level | Get the current level of the digital reference input
*FastResourceAPI* | [**ApiAudioRoomInUseActivityLevelGet**](docs/FastResourceAPI.md#apiaudioroominuseactivitylevelget) | **Get** /api/audio/roomInUse/activityLevel | Get the current room in use activity level
*FirmwareUpdateAPI* | [**ApiFirmwareUpdateStateGet**](docs/FirmwareUpdateAPI.md#apifirmwareupdatestateget) | **Get** /api/firmware/update/state | Get the state of a firmware update
*SSCAPI* | [**ApiSscSchemaGet**](docs/SSCAPI.md#apisscschemaget) | **Get** /api/ssc/schema | Get the address tree
*SSCAPI* | [**ApiSscStateSubscriptionsGet**](docs/SSCAPI.md#apisscstatesubscriptionsget) | **Get** /api/ssc/state/subscriptions | Start a subscription
*SSCAPI* | [**ApiSscStateSubscriptionsSessionUUIDAddPut**](docs/SSCAPI.md#apisscstatesubscriptionssessionuuidaddput) | **Put** /api/ssc/state/subscriptions/{sessionUUID}/add | Add resource(s) to the subscription list
*SSCAPI* | [**ApiSscStateSubscriptionsSessionUUIDDelete**](docs/SSCAPI.md#apisscstatesubscriptionssessionuuiddelete) | **Delete** /api/ssc/state/subscriptions/{sessionUUID} | End an existing subscription
*SSCAPI* | [**ApiSscStateSubscriptionsSessionUUIDGet**](docs/SSCAPI.md#apisscstatesubscriptionssessionuuidget) | **Get** /api/ssc/state/subscriptions/{sessionUUID} | Get the subscription list
*SSCAPI* | [**ApiSscStateSubscriptionsSessionUUIDPut**](docs/SSCAPI.md#apisscstatesubscriptionssessionuuidput) | **Put** /api/ssc/state/subscriptions/{sessionUUID} | Set or change the list of subscriptions associated with the sessionUUID
*SSCAPI* | [**ApiSscStateSubscriptionsSessionUUIDRemovePut**](docs/SSCAPI.md#apisscstatesubscriptionssessionuuidremoveput) | **Put** /api/ssc/state/subscriptions/{sessionUUID}/remove | Remove resource(s) from the subscription list
*SSCAPI* | [**ApiSscVersionGet**](docs/SSCAPI.md#apisscversionget) | **Get** /api/ssc/version | Get the schema version


## Documentation For Models

 - [ApiAudioEqualizerGet200Response](docs/ApiAudioEqualizerGet200Response.md)
 - [ApiAudioEqualizerPutRequest](docs/ApiAudioEqualizerPutRequest.md)
 - [ApiAudioInputsDanteReferenceGet200Response](docs/ApiAudioInputsDanteReferenceGet200Response.md)
 - [ApiAudioInputsDanteReferencePutRequest](docs/ApiAudioInputsDanteReferencePutRequest.md)
 - [ApiAudioInputsMicrophoneBeamDirectionGet200Response](docs/ApiAudioInputsMicrophoneBeamDirectionGet200Response.md)
 - [ApiAudioInputsMicrophoneBeamGet200Response](docs/ApiAudioInputsMicrophoneBeamGet200Response.md)
 - [ApiAudioInputsMicrophoneBeamPutRequest](docs/ApiAudioInputsMicrophoneBeamPutRequest.md)
 - [ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner](docs/ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInner.md)
 - [ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth](docs/ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth.md)
 - [ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerElevation](docs/ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerElevation.md)
 - [ApiAudioInputsMicrophoneExclusionZonesIdGet200Response](docs/ApiAudioInputsMicrophoneExclusionZonesIdGet200Response.md)
 - [ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth](docs/ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseAzimuth.md)
 - [ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation](docs/ApiAudioInputsMicrophoneExclusionZonesIdGet200ResponseElevation.md)
 - [ApiAudioInputsMicrophoneExclusionZonesIdPutRequest](docs/ApiAudioInputsMicrophoneExclusionZonesIdPutRequest.md)
 - [ApiAudioInputsMicrophoneLevelGet200Response](docs/ApiAudioInputsMicrophoneLevelGet200Response.md)
 - [ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner](docs/ApiAudioInputsMicrophonePriorityZonesGet200ResponseInner.md)
 - [ApiAudioInputsMicrophonePriorityZonesIdGet200Response](docs/ApiAudioInputsMicrophonePriorityZonesIdGet200Response.md)
 - [ApiAudioInputsMicrophonePriorityZonesIdPutRequest](docs/ApiAudioInputsMicrophonePriorityZonesIdPutRequest.md)
 - [ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth](docs/ApiAudioInputsMicrophonePriorityZonesIdPutRequestAzimuth.md)
 - [ApiAudioInputsMicrophonePriorityZonesIdPutRequestElevation](docs/ApiAudioInputsMicrophonePriorityZonesIdPutRequestElevation.md)
 - [ApiAudioInputsReferenceLevelGet200Response](docs/ApiAudioInputsReferenceLevelGet200Response.md)
 - [ApiAudioNoiseGateGet200Response](docs/ApiAudioNoiseGateGet200Response.md)
 - [ApiAudioNoiseGatePutRequest](docs/ApiAudioNoiseGatePutRequest.md)
 - [ApiAudioOutputsAnalogGet200Response](docs/ApiAudioOutputsAnalogGet200Response.md)
 - [ApiAudioOutputsAnalogPutRequest](docs/ApiAudioOutputsAnalogPutRequest.md)
 - [ApiAudioOutputsDanteFarEndGet200Response](docs/ApiAudioOutputsDanteFarEndGet200Response.md)
 - [ApiAudioOutputsDanteFarEndPutRequest](docs/ApiAudioOutputsDanteFarEndPutRequest.md)
 - [ApiAudioOutputsDanteLocalGet200Response](docs/ApiAudioOutputsDanteLocalGet200Response.md)
 - [ApiAudioOutputsDanteLocalPutRequest](docs/ApiAudioOutputsDanteLocalPutRequest.md)
 - [ApiAudioOutputsGlobalMuteGet200Response](docs/ApiAudioOutputsGlobalMuteGet200Response.md)
 - [ApiAudioOutputsGlobalMutePutRequest](docs/ApiAudioOutputsGlobalMutePutRequest.md)
 - [ApiAudioRoomInUseActivityLevelGet200Response](docs/ApiAudioRoomInUseActivityLevelGet200Response.md)
 - [ApiAudioRoomInUseConfigGet200Response](docs/ApiAudioRoomInUseConfigGet200Response.md)
 - [ApiAudioRoomInUseGet200Response](docs/ApiAudioRoomInUseGet200Response.md)
 - [ApiAudioVoiceLiftGet200Response](docs/ApiAudioVoiceLiftGet200Response.md)
 - [ApiAudioVoiceLiftPutRequest](docs/ApiAudioVoiceLiftPutRequest.md)
 - [ApiDeviceIdentificationGet200Response](docs/ApiDeviceIdentificationGet200Response.md)
 - [ApiDeviceIdentificationPutRequest](docs/ApiDeviceIdentificationPutRequest.md)
 - [ApiDeviceIdentityGet200Response](docs/ApiDeviceIdentityGet200Response.md)
 - [ApiDeviceLedsRingGet200Response](docs/ApiDeviceLedsRingGet200Response.md)
 - [ApiDeviceLedsRingGet200ResponseMicCustom](docs/ApiDeviceLedsRingGet200ResponseMicCustom.md)
 - [ApiDeviceLedsRingGet200ResponseMicOn](docs/ApiDeviceLedsRingGet200ResponseMicOn.md)
 - [ApiDeviceLedsRingPutRequest](docs/ApiDeviceLedsRingPutRequest.md)
 - [ApiDeviceLedsRingPutRequestMicCustom](docs/ApiDeviceLedsRingPutRequestMicCustom.md)
 - [ApiDeviceLedsRingPutRequestMicOn](docs/ApiDeviceLedsRingPutRequestMicOn.md)
 - [ApiDevicePowerPoeDaisychainGet200Response](docs/ApiDevicePowerPoeDaisychainGet200Response.md)
 - [ApiDeviceSiteGet200Response](docs/ApiDeviceSiteGet200Response.md)
 - [ApiDeviceStateGet200Response](docs/ApiDeviceStateGet200Response.md)
 - [ApiFirmwareUpdateStateGet200Response](docs/ApiFirmwareUpdateStateGet200Response.md)
 - [ApiSscStateSubscriptionsSessionUUIDAddPut400Response](docs/ApiSscStateSubscriptionsSessionUUIDAddPut400Response.md)
 - [ApiSscStateSubscriptionsSessionUUIDRemovePut400Response](docs/ApiSscStateSubscriptionsSessionUUIDRemovePut400Response.md)
 - [ApiSscVersionGet200Response](docs/ApiSscVersionGet200Response.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuthApi

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



