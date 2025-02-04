// Code generated by "stringer -type=MemoryKind -trimprefix=MemoryKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MemoryKindUnknown-0]
	_ = x[MemoryKindAddressValue-1]
	_ = x[MemoryKindStringValue-2]
	_ = x[MemoryKindCharacterValue-3]
	_ = x[MemoryKindNumberValue-4]
	_ = x[MemoryKindArrayValueBase-5]
	_ = x[MemoryKindDictionaryValueBase-6]
	_ = x[MemoryKindCompositeValueBase-7]
	_ = x[MemoryKindSimpleCompositeValueBase-8]
	_ = x[MemoryKindOptionalValue-9]
	_ = x[MemoryKindTypeValue-10]
	_ = x[MemoryKindPathValue-11]
	_ = x[MemoryKindStorageCapabilityValue-12]
	_ = x[MemoryKindPathLinkValue-13]
	_ = x[MemoryKindAccountLinkValue-14]
	_ = x[MemoryKindStorageReferenceValue-15]
	_ = x[MemoryKindAccountReferenceValue-16]
	_ = x[MemoryKindEphemeralReferenceValue-17]
	_ = x[MemoryKindInterpretedFunctionValue-18]
	_ = x[MemoryKindHostFunctionValue-19]
	_ = x[MemoryKindBoundFunctionValue-20]
	_ = x[MemoryKindBigInt-21]
	_ = x[MemoryKindSimpleCompositeValue-22]
	_ = x[MemoryKindPublishedValue-23]
	_ = x[MemoryKindAtreeArrayDataSlab-24]
	_ = x[MemoryKindAtreeArrayMetaDataSlab-25]
	_ = x[MemoryKindAtreeArrayElementOverhead-26]
	_ = x[MemoryKindAtreeMapDataSlab-27]
	_ = x[MemoryKindAtreeMapMetaDataSlab-28]
	_ = x[MemoryKindAtreeMapElementOverhead-29]
	_ = x[MemoryKindAtreeMapPreAllocatedElement-30]
	_ = x[MemoryKindAtreeEncodedSlab-31]
	_ = x[MemoryKindPrimitiveStaticType-32]
	_ = x[MemoryKindCompositeStaticType-33]
	_ = x[MemoryKindInterfaceStaticType-34]
	_ = x[MemoryKindVariableSizedStaticType-35]
	_ = x[MemoryKindConstantSizedStaticType-36]
	_ = x[MemoryKindDictionaryStaticType-37]
	_ = x[MemoryKindOptionalStaticType-38]
	_ = x[MemoryKindRestrictedStaticType-39]
	_ = x[MemoryKindReferenceStaticType-40]
	_ = x[MemoryKindCapabilityStaticType-41]
	_ = x[MemoryKindFunctionStaticType-42]
	_ = x[MemoryKindCadenceVoidValue-43]
	_ = x[MemoryKindCadenceOptionalValue-44]
	_ = x[MemoryKindCadenceBoolValue-45]
	_ = x[MemoryKindCadenceStringValue-46]
	_ = x[MemoryKindCadenceCharacterValue-47]
	_ = x[MemoryKindCadenceAddressValue-48]
	_ = x[MemoryKindCadenceIntValue-49]
	_ = x[MemoryKindCadenceNumberValue-50]
	_ = x[MemoryKindCadenceArrayValueBase-51]
	_ = x[MemoryKindCadenceArrayValueLength-52]
	_ = x[MemoryKindCadenceDictionaryValue-53]
	_ = x[MemoryKindCadenceKeyValuePair-54]
	_ = x[MemoryKindCadenceStructValueBase-55]
	_ = x[MemoryKindCadenceStructValueSize-56]
	_ = x[MemoryKindCadenceResourceValueBase-57]
	_ = x[MemoryKindCadenceResourceValueSize-58]
	_ = x[MemoryKindCadenceEventValueBase-59]
	_ = x[MemoryKindCadenceEventValueSize-60]
	_ = x[MemoryKindCadenceContractValueBase-61]
	_ = x[MemoryKindCadenceContractValueSize-62]
	_ = x[MemoryKindCadenceEnumValueBase-63]
	_ = x[MemoryKindCadenceEnumValueSize-64]
	_ = x[MemoryKindCadencePathLinkValue-65]
	_ = x[MemoryKindCadencePathValue-66]
	_ = x[MemoryKindCadenceTypeValue-67]
	_ = x[MemoryKindCadenceStorageCapabilityValue-68]
	_ = x[MemoryKindCadenceFunctionValue-69]
	_ = x[MemoryKindCadenceSimpleType-70]
	_ = x[MemoryKindCadenceOptionalType-71]
	_ = x[MemoryKindCadenceVariableSizedArrayType-72]
	_ = x[MemoryKindCadenceConstantSizedArrayType-73]
	_ = x[MemoryKindCadenceDictionaryType-74]
	_ = x[MemoryKindCadenceField-75]
	_ = x[MemoryKindCadenceParameter-76]
	_ = x[MemoryKindCadenceStructType-77]
	_ = x[MemoryKindCadenceResourceType-78]
	_ = x[MemoryKindCadenceEventType-79]
	_ = x[MemoryKindCadenceContractType-80]
	_ = x[MemoryKindCadenceStructInterfaceType-81]
	_ = x[MemoryKindCadenceResourceInterfaceType-82]
	_ = x[MemoryKindCadenceContractInterfaceType-83]
	_ = x[MemoryKindCadenceFunctionType-84]
	_ = x[MemoryKindCadenceReferenceType-85]
	_ = x[MemoryKindCadenceRestrictedType-86]
	_ = x[MemoryKindCadenceCapabilityType-87]
	_ = x[MemoryKindCadenceEnumType-88]
	_ = x[MemoryKindRawString-89]
	_ = x[MemoryKindAddressLocation-90]
	_ = x[MemoryKindBytes-91]
	_ = x[MemoryKindVariable-92]
	_ = x[MemoryKindCompositeTypeInfo-93]
	_ = x[MemoryKindCompositeField-94]
	_ = x[MemoryKindInvocation-95]
	_ = x[MemoryKindStorageMap-96]
	_ = x[MemoryKindStorageKey-97]
	_ = x[MemoryKindTypeToken-98]
	_ = x[MemoryKindErrorToken-99]
	_ = x[MemoryKindSpaceToken-100]
	_ = x[MemoryKindProgram-101]
	_ = x[MemoryKindIdentifier-102]
	_ = x[MemoryKindArgument-103]
	_ = x[MemoryKindBlock-104]
	_ = x[MemoryKindFunctionBlock-105]
	_ = x[MemoryKindParameter-106]
	_ = x[MemoryKindParameterList-107]
	_ = x[MemoryKindTransfer-108]
	_ = x[MemoryKindMembers-109]
	_ = x[MemoryKindTypeAnnotation-110]
	_ = x[MemoryKindDictionaryEntry-111]
	_ = x[MemoryKindFunctionDeclaration-112]
	_ = x[MemoryKindCompositeDeclaration-113]
	_ = x[MemoryKindInterfaceDeclaration-114]
	_ = x[MemoryKindEnumCaseDeclaration-115]
	_ = x[MemoryKindFieldDeclaration-116]
	_ = x[MemoryKindTransactionDeclaration-117]
	_ = x[MemoryKindImportDeclaration-118]
	_ = x[MemoryKindVariableDeclaration-119]
	_ = x[MemoryKindSpecialFunctionDeclaration-120]
	_ = x[MemoryKindPragmaDeclaration-121]
	_ = x[MemoryKindAssignmentStatement-122]
	_ = x[MemoryKindBreakStatement-123]
	_ = x[MemoryKindContinueStatement-124]
	_ = x[MemoryKindEmitStatement-125]
	_ = x[MemoryKindExpressionStatement-126]
	_ = x[MemoryKindForStatement-127]
	_ = x[MemoryKindIfStatement-128]
	_ = x[MemoryKindReturnStatement-129]
	_ = x[MemoryKindSwapStatement-130]
	_ = x[MemoryKindSwitchStatement-131]
	_ = x[MemoryKindWhileStatement-132]
	_ = x[MemoryKindBooleanExpression-133]
	_ = x[MemoryKindVoidExpression-134]
	_ = x[MemoryKindNilExpression-135]
	_ = x[MemoryKindStringExpression-136]
	_ = x[MemoryKindIntegerExpression-137]
	_ = x[MemoryKindFixedPointExpression-138]
	_ = x[MemoryKindArrayExpression-139]
	_ = x[MemoryKindDictionaryExpression-140]
	_ = x[MemoryKindIdentifierExpression-141]
	_ = x[MemoryKindInvocationExpression-142]
	_ = x[MemoryKindMemberExpression-143]
	_ = x[MemoryKindIndexExpression-144]
	_ = x[MemoryKindConditionalExpression-145]
	_ = x[MemoryKindUnaryExpression-146]
	_ = x[MemoryKindBinaryExpression-147]
	_ = x[MemoryKindFunctionExpression-148]
	_ = x[MemoryKindCastingExpression-149]
	_ = x[MemoryKindCreateExpression-150]
	_ = x[MemoryKindDestroyExpression-151]
	_ = x[MemoryKindReferenceExpression-152]
	_ = x[MemoryKindForceExpression-153]
	_ = x[MemoryKindPathExpression-154]
	_ = x[MemoryKindConstantSizedType-155]
	_ = x[MemoryKindDictionaryType-156]
	_ = x[MemoryKindFunctionType-157]
	_ = x[MemoryKindInstantiationType-158]
	_ = x[MemoryKindNominalType-159]
	_ = x[MemoryKindOptionalType-160]
	_ = x[MemoryKindReferenceType-161]
	_ = x[MemoryKindRestrictedType-162]
	_ = x[MemoryKindVariableSizedType-163]
	_ = x[MemoryKindPosition-164]
	_ = x[MemoryKindRange-165]
	_ = x[MemoryKindElaboration-166]
	_ = x[MemoryKindActivation-167]
	_ = x[MemoryKindActivationEntries-168]
	_ = x[MemoryKindVariableSizedSemaType-169]
	_ = x[MemoryKindConstantSizedSemaType-170]
	_ = x[MemoryKindDictionarySemaType-171]
	_ = x[MemoryKindOptionalSemaType-172]
	_ = x[MemoryKindRestrictedSemaType-173]
	_ = x[MemoryKindReferenceSemaType-174]
	_ = x[MemoryKindCapabilitySemaType-175]
	_ = x[MemoryKindOrderedMap-176]
	_ = x[MemoryKindOrderedMapEntryList-177]
	_ = x[MemoryKindOrderedMapEntry-178]
	_ = x[MemoryKindLast-179]
}

const _MemoryKind_name = "UnknownAddressValueStringValueCharacterValueNumberValueArrayValueBaseDictionaryValueBaseCompositeValueBaseSimpleCompositeValueBaseOptionalValueTypeValuePathValueStorageCapabilityValuePathLinkValueAccountLinkValueStorageReferenceValueAccountReferenceValueEphemeralReferenceValueInterpretedFunctionValueHostFunctionValueBoundFunctionValueBigIntSimpleCompositeValuePublishedValueAtreeArrayDataSlabAtreeArrayMetaDataSlabAtreeArrayElementOverheadAtreeMapDataSlabAtreeMapMetaDataSlabAtreeMapElementOverheadAtreeMapPreAllocatedElementAtreeEncodedSlabPrimitiveStaticTypeCompositeStaticTypeInterfaceStaticTypeVariableSizedStaticTypeConstantSizedStaticTypeDictionaryStaticTypeOptionalStaticTypeRestrictedStaticTypeReferenceStaticTypeCapabilityStaticTypeFunctionStaticTypeCadenceVoidValueCadenceOptionalValueCadenceBoolValueCadenceStringValueCadenceCharacterValueCadenceAddressValueCadenceIntValueCadenceNumberValueCadenceArrayValueBaseCadenceArrayValueLengthCadenceDictionaryValueCadenceKeyValuePairCadenceStructValueBaseCadenceStructValueSizeCadenceResourceValueBaseCadenceResourceValueSizeCadenceEventValueBaseCadenceEventValueSizeCadenceContractValueBaseCadenceContractValueSizeCadenceEnumValueBaseCadenceEnumValueSizeCadencePathLinkValueCadencePathValueCadenceTypeValueCadenceStorageCapabilityValueCadenceFunctionValueCadenceSimpleTypeCadenceOptionalTypeCadenceVariableSizedArrayTypeCadenceConstantSizedArrayTypeCadenceDictionaryTypeCadenceFieldCadenceParameterCadenceStructTypeCadenceResourceTypeCadenceEventTypeCadenceContractTypeCadenceStructInterfaceTypeCadenceResourceInterfaceTypeCadenceContractInterfaceTypeCadenceFunctionTypeCadenceReferenceTypeCadenceRestrictedTypeCadenceCapabilityTypeCadenceEnumTypeRawStringAddressLocationBytesVariableCompositeTypeInfoCompositeFieldInvocationStorageMapStorageKeyTypeTokenErrorTokenSpaceTokenProgramIdentifierArgumentBlockFunctionBlockParameterParameterListTransferMembersTypeAnnotationDictionaryEntryFunctionDeclarationCompositeDeclarationInterfaceDeclarationEnumCaseDeclarationFieldDeclarationTransactionDeclarationImportDeclarationVariableDeclarationSpecialFunctionDeclarationPragmaDeclarationAssignmentStatementBreakStatementContinueStatementEmitStatementExpressionStatementForStatementIfStatementReturnStatementSwapStatementSwitchStatementWhileStatementBooleanExpressionVoidExpressionNilExpressionStringExpressionIntegerExpressionFixedPointExpressionArrayExpressionDictionaryExpressionIdentifierExpressionInvocationExpressionMemberExpressionIndexExpressionConditionalExpressionUnaryExpressionBinaryExpressionFunctionExpressionCastingExpressionCreateExpressionDestroyExpressionReferenceExpressionForceExpressionPathExpressionConstantSizedTypeDictionaryTypeFunctionTypeInstantiationTypeNominalTypeOptionalTypeReferenceTypeRestrictedTypeVariableSizedTypePositionRangeElaborationActivationActivationEntriesVariableSizedSemaTypeConstantSizedSemaTypeDictionarySemaTypeOptionalSemaTypeRestrictedSemaTypeReferenceSemaTypeCapabilitySemaTypeOrderedMapOrderedMapEntryListOrderedMapEntryLast"

var _MemoryKind_index = [...]uint16{0, 7, 19, 30, 44, 55, 69, 88, 106, 130, 143, 152, 161, 183, 196, 212, 233, 254, 277, 301, 318, 336, 342, 362, 376, 394, 416, 441, 457, 477, 500, 527, 543, 562, 581, 600, 623, 646, 666, 684, 704, 723, 743, 761, 777, 797, 813, 831, 852, 871, 886, 904, 925, 948, 970, 989, 1011, 1033, 1057, 1081, 1102, 1123, 1147, 1171, 1191, 1211, 1231, 1247, 1263, 1292, 1312, 1329, 1348, 1377, 1406, 1427, 1439, 1455, 1472, 1491, 1507, 1526, 1552, 1580, 1608, 1627, 1647, 1668, 1689, 1704, 1713, 1728, 1733, 1741, 1758, 1772, 1782, 1792, 1802, 1811, 1821, 1831, 1838, 1848, 1856, 1861, 1874, 1883, 1896, 1904, 1911, 1925, 1940, 1959, 1979, 1999, 2018, 2034, 2056, 2073, 2092, 2118, 2135, 2154, 2168, 2185, 2198, 2217, 2229, 2240, 2255, 2268, 2283, 2297, 2314, 2328, 2341, 2357, 2374, 2394, 2409, 2429, 2449, 2469, 2485, 2500, 2521, 2536, 2552, 2570, 2587, 2603, 2620, 2639, 2654, 2668, 2685, 2699, 2711, 2728, 2739, 2751, 2764, 2778, 2795, 2803, 2808, 2819, 2829, 2846, 2867, 2888, 2906, 2922, 2940, 2957, 2975, 2985, 3004, 3019, 3023}

func (i MemoryKind) String() string {
	if i >= MemoryKind(len(_MemoryKind_index)-1) {
		return "MemoryKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MemoryKind_name[_MemoryKind_index[i]:_MemoryKind_index[i+1]]
}
