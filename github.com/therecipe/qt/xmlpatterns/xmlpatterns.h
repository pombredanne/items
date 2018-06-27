// +build !minimal

#pragma once

#ifndef GO_QTXMLPATTERNS_H
#define GO_QTXMLPATTERNS_H

#include <stdint.h>

#ifdef __cplusplus
int QAbstractMessageHandler_QAbstractMessageHandler_QRegisterMetaType();
int QAbstractUriResolver_QAbstractUriResolver_QRegisterMetaType();
extern "C" {
#endif

struct QtXmlPatterns_PackedString { char* data; long long len; };
struct QtXmlPatterns_PackedList { void* data; long long len; };
void QAbstractMessageHandler_DestroyQAbstractMessageHandler(void* ptr);
void QAbstractMessageHandler_DestroyQAbstractMessageHandlerDefault(void* ptr);
void* QAbstractMessageHandler___dynamicPropertyNames_atList(void* ptr, int i, void* p);
void QAbstractMessageHandler___dynamicPropertyNames_setList(void* ptr, void* i, void* p);
void* QAbstractMessageHandler___dynamicPropertyNames_newList(void* ptr);
void* QAbstractMessageHandler___findChildren_atList2(void* ptr, int i, void* p);
void QAbstractMessageHandler___findChildren_setList2(void* ptr, void* i, void* p);
void* QAbstractMessageHandler___findChildren_newList2(void* ptr);
void* QAbstractMessageHandler___findChildren_atList3(void* ptr, int i, void* p);
void QAbstractMessageHandler___findChildren_setList3(void* ptr, void* i, void* p);
void* QAbstractMessageHandler___findChildren_newList3(void* ptr);
void* QAbstractMessageHandler___findChildren_atList(void* ptr, int i, void* p);
void QAbstractMessageHandler___findChildren_setList(void* ptr, void* i, void* p);
void* QAbstractMessageHandler___findChildren_newList(void* ptr);
void* QAbstractMessageHandler___children_atList(void* ptr, int i, void* p);
void QAbstractMessageHandler___children_setList(void* ptr, void* i, void* p);
void* QAbstractMessageHandler___children_newList(void* ptr);
char QAbstractMessageHandler_EventDefault(void* ptr, void* e);
char QAbstractMessageHandler_EventFilterDefault(void* ptr, void* watched, void* event);
void QAbstractMessageHandler_ChildEventDefault(void* ptr, void* event);
void QAbstractMessageHandler_ConnectNotifyDefault(void* ptr, void* sign);
void QAbstractMessageHandler_CustomEventDefault(void* ptr, void* event);
void QAbstractMessageHandler_DeleteLaterDefault(void* ptr);
void QAbstractMessageHandler_DisconnectNotifyDefault(void* ptr, void* sign);
void QAbstractMessageHandler_TimerEventDefault(void* ptr, void* event);
void* QAbstractMessageHandler_MetaObjectDefault(void* ptr);
void* QAbstractUriResolver_NewQAbstractUriResolver(void* parent);
void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr);
void QAbstractUriResolver_DestroyQAbstractUriResolverDefault(void* ptr);
void* QAbstractUriResolver_Resolve(void* ptr, void* relative, void* baseURI);
void* QAbstractUriResolver___dynamicPropertyNames_atList(void* ptr, int i, void* p);
void QAbstractUriResolver___dynamicPropertyNames_setList(void* ptr, void* i, void* p);
void* QAbstractUriResolver___dynamicPropertyNames_newList(void* ptr);
void* QAbstractUriResolver___findChildren_atList2(void* ptr, int i, void* p);
void QAbstractUriResolver___findChildren_setList2(void* ptr, void* i, void* p);
void* QAbstractUriResolver___findChildren_newList2(void* ptr);
void* QAbstractUriResolver___findChildren_atList3(void* ptr, int i, void* p);
void QAbstractUriResolver___findChildren_setList3(void* ptr, void* i, void* p);
void* QAbstractUriResolver___findChildren_newList3(void* ptr);
void* QAbstractUriResolver___findChildren_atList(void* ptr, int i, void* p);
void QAbstractUriResolver___findChildren_setList(void* ptr, void* i, void* p);
void* QAbstractUriResolver___findChildren_newList(void* ptr);
void* QAbstractUriResolver___children_atList(void* ptr, int i, void* p);
void QAbstractUriResolver___children_setList(void* ptr, void* i, void* p);
void* QAbstractUriResolver___children_newList(void* ptr);
char QAbstractUriResolver_EventDefault(void* ptr, void* e);
char QAbstractUriResolver_EventFilterDefault(void* ptr, void* watched, void* event);
void QAbstractUriResolver_ChildEventDefault(void* ptr, void* event);
void QAbstractUriResolver_ConnectNotifyDefault(void* ptr, void* sign);
void QAbstractUriResolver_CustomEventDefault(void* ptr, void* event);
void QAbstractUriResolver_DeleteLaterDefault(void* ptr);
void QAbstractUriResolver_DisconnectNotifyDefault(void* ptr, void* sign);
void QAbstractUriResolver_TimerEventDefault(void* ptr, void* event);
void* QAbstractUriResolver_MetaObjectDefault(void* ptr);
void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr);
void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModelDefault(void* ptr);
void* QAbstractXmlNodeModel_SourceLocation(void* ptr, void* index);
struct QtXmlPatterns_PackedString QAbstractXmlNodeModel_StringValue(void* ptr, void* n);
void* QAbstractXmlNodeModel_BaseUri(void* ptr, void* n);
void* QAbstractXmlNodeModel_DocumentUri(void* ptr, void* n);
void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node);
struct QtXmlPatterns_PackedList QAbstractXmlNodeModel_NamespaceBindings(void* ptr, void* n);
struct QtXmlPatterns_PackedList QAbstractXmlNodeModel_NodesByIdref(void* ptr, void* idref);
void* QAbstractXmlNodeModel_Name(void* ptr, void* ni);
void* QAbstractXmlNodeModel_CreateIndex(void* ptr, long long data);
void* QAbstractXmlNodeModel_CreateIndex3(void* ptr, long long data, long long additionalData);
void* QAbstractXmlNodeModel_CreateIndex2(void* ptr, void* pointer, long long additionalData);
void* QAbstractXmlNodeModel_ElementById(void* ptr, void* id);
void* QAbstractXmlNodeModel_NextFromSimpleAxis(void* ptr, long long axis, void* origin);
void* QAbstractXmlNodeModel_Root(void* ptr, void* n);
long long QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2);
long long QAbstractXmlNodeModel_Kind(void* ptr, void* ni);
void* QAbstractXmlNodeModel___namespaceBindings_atList(void* ptr, int i, void* p);
void QAbstractXmlNodeModel___namespaceBindings_setList(void* ptr, void* i, void* p);
void* QAbstractXmlNodeModel___namespaceBindings_newList(void* ptr);
void* QAbstractXmlNodeModel___nodesByIdref_atList(void* ptr, int i, void* p);
void QAbstractXmlNodeModel___nodesByIdref_setList(void* ptr, void* i, void* p);
void* QAbstractXmlNodeModel___nodesByIdref_newList(void* ptr);
void* QAbstractXmlReceiver_NewQAbstractXmlReceiver();
void QAbstractXmlReceiver_AtomicValue(void* ptr, void* value);
void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value);
void QAbstractXmlReceiver_Characters(void* ptr, void* value);
void QAbstractXmlReceiver_Comment(void* ptr, struct QtXmlPatterns_PackedString value);
void QAbstractXmlReceiver_EndDocument(void* ptr);
void QAbstractXmlReceiver_EndElement(void* ptr);
void QAbstractXmlReceiver_EndOfSequence(void* ptr);
void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name);
void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, struct QtXmlPatterns_PackedString value);
void QAbstractXmlReceiver_StartDocument(void* ptr);
void QAbstractXmlReceiver_StartElement(void* ptr, void* name);
void QAbstractXmlReceiver_StartOfSequence(void* ptr);
void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr);
void QAbstractXmlReceiver_DestroyQAbstractXmlReceiverDefault(void* ptr);
void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr);
void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModelDefault(void* ptr);
struct QtXmlPatterns_PackedString QSimpleXmlNodeModel_StringValue(void* ptr, void* node);
struct QtXmlPatterns_PackedString QSimpleXmlNodeModel_StringValueDefault(void* ptr, void* node);
void* QSimpleXmlNodeModel_BaseUri(void* ptr, void* node);
void* QSimpleXmlNodeModel_BaseUriDefault(void* ptr, void* node);
struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NamespaceBindings(void* ptr, void* node);
struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NamespaceBindingsDefault(void* ptr, void* node);
struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NodesByIdref(void* ptr, void* idref);
struct QtXmlPatterns_PackedList QSimpleXmlNodeModel_NodesByIdrefDefault(void* ptr, void* idref);
void* QSimpleXmlNodeModel_NamePool(void* ptr);
void* QSimpleXmlNodeModel_ElementById(void* ptr, void* id);
void* QSimpleXmlNodeModel_ElementByIdDefault(void* ptr, void* id);
void* QSimpleXmlNodeModel_DocumentUri(void* ptr, void* n);
void* QSimpleXmlNodeModel_DocumentUriDefault(void* ptr, void* n);
void* QSimpleXmlNodeModel_TypedValue(void* ptr, void* node);
void* QSimpleXmlNodeModel_TypedValueDefault(void* ptr, void* node);
void* QSimpleXmlNodeModel_Name(void* ptr, void* ni);
void* QSimpleXmlNodeModel_NameDefault(void* ptr, void* ni);
void* QSimpleXmlNodeModel_NextFromSimpleAxis(void* ptr, long long axis, void* origin);
void* QSimpleXmlNodeModel_NextFromSimpleAxisDefault(void* ptr, long long axis, void* origin);
void* QSimpleXmlNodeModel_Root(void* ptr, void* n);
void* QSimpleXmlNodeModel_RootDefault(void* ptr, void* n);
long long QSimpleXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2);
long long QSimpleXmlNodeModel_CompareOrderDefault(void* ptr, void* ni1, void* ni2);
long long QSimpleXmlNodeModel_Kind(void* ptr, void* ni);
long long QSimpleXmlNodeModel_KindDefault(void* ptr, void* ni);
void* QSourceLocation_NewQSourceLocation();
void* QSourceLocation_NewQSourceLocation2(void* other);
void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c);
void QSourceLocation_SetColumn(void* ptr, long long newColumn);
void QSourceLocation_SetLine(void* ptr, long long newLine);
void QSourceLocation_SetUri(void* ptr, void* newUri);
void QSourceLocation_DestroyQSourceLocation(void* ptr);
void* QSourceLocation_Uri(void* ptr);
char QSourceLocation_IsNull(void* ptr);
long long QSourceLocation_Column(void* ptr);
long long QSourceLocation_Line(void* ptr);
void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice);
void QXmlFormatter_SetIndentationDepth(void* ptr, int depth);
int QXmlFormatter_IndentationDepth(void* ptr);
void* QXmlItem_NewQXmlItem();
void* QXmlItem_NewQXmlItem4(void* atomicValue);
void* QXmlItem_NewQXmlItem2(void* other);
void* QXmlItem_NewQXmlItem3(void* node);
void QXmlItem_DestroyQXmlItem(void* ptr);
void* QXmlItem_ToAtomicValue(void* ptr);
void* QXmlItem_ToNodeModelIndex(void* ptr);
char QXmlItem_IsAtomicValue(void* ptr);
char QXmlItem_IsNode(void* ptr);
char QXmlItem_IsNull(void* ptr);
void* QXmlName_QXmlName_FromClarkName(struct QtXmlPatterns_PackedString clarkName, void* namePool);
void* QXmlName_NewQXmlName();
void* QXmlName_NewQXmlName2(void* namePool, struct QtXmlPatterns_PackedString localName, struct QtXmlPatterns_PackedString namespaceURI, struct QtXmlPatterns_PackedString prefix);
void* QXmlName_NewQXmlName3(void* other);
char QXmlName_QXmlName_IsNCName(struct QtXmlPatterns_PackedString candidate);
struct QtXmlPatterns_PackedString QXmlName_LocalName(void* ptr, void* namePool);
struct QtXmlPatterns_PackedString QXmlName_NamespaceUri(void* ptr, void* namePool);
struct QtXmlPatterns_PackedString QXmlName_Prefix(void* ptr, void* namePool);
struct QtXmlPatterns_PackedString QXmlName_ToClarkName(void* ptr, void* namePool);
char QXmlName_IsNull(void* ptr);
void* QXmlNamePool_NewQXmlNamePool();
void* QXmlNamePool_NewQXmlNamePool2(void* other);
void QXmlNamePool_DestroyQXmlNamePool(void* ptr);
void* QXmlNodeModelIndex_NewQXmlNodeModelIndex();
void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other);
char QXmlNodeModelIndex_IsNull(void* ptr);
void* QXmlNodeModelIndex_Model(void* ptr);
long long QXmlNodeModelIndex_AdditionalData(void* ptr);
long long QXmlNodeModelIndex_Data(void* ptr);
void* QXmlNodeModelIndex_InternalPointer(void* ptr);
void* QXmlNodeModelIndex___namespaceBindings_atList(void* ptr, int i, void* p);
void QXmlNodeModelIndex___namespaceBindings_setList(void* ptr, void* i, void* p);
void* QXmlNodeModelIndex___namespaceBindings_newList(void* ptr);
void* QXmlQuery_NewQXmlQuery();
void* QXmlQuery_NewQXmlQuery4(long long queryLanguage, void* np);
void* QXmlQuery_NewQXmlQuery3(void* np);
void* QXmlQuery_NewQXmlQuery2(void* other);
char QXmlQuery_SetFocus3(void* ptr, void* document);
char QXmlQuery_SetFocus4(void* ptr, struct QtXmlPatterns_PackedString focus);
char QXmlQuery_SetFocus2(void* ptr, void* documentURI);
void QXmlQuery_BindVariable4(void* ptr, struct QtXmlPatterns_PackedString localName, void* device);
void QXmlQuery_BindVariable2(void* ptr, struct QtXmlPatterns_PackedString localName, void* value);
void QXmlQuery_BindVariable6(void* ptr, struct QtXmlPatterns_PackedString localName, void* query);
void QXmlQuery_BindVariable3(void* ptr, void* name, void* device);
void QXmlQuery_BindVariable(void* ptr, void* name, void* value);
void QXmlQuery_BindVariable5(void* ptr, void* name, void* query);
void QXmlQuery_SetFocus(void* ptr, void* item);
void QXmlQuery_SetInitialTemplateName2(void* ptr, struct QtXmlPatterns_PackedString localName);
void QXmlQuery_SetInitialTemplateName(void* ptr, void* name);
void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler);
void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager);
void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI);
void QXmlQuery_SetQuery2(void* ptr, struct QtXmlPatterns_PackedString sourceCode, void* documentURI);
void QXmlQuery_SetQuery3(void* ptr, void* queryURI, void* baseURI);
void QXmlQuery_SetUriResolver(void* ptr, void* resolver);
void QXmlQuery_DestroyQXmlQuery(void* ptr);
void* QXmlQuery_MessageHandler(void* ptr);
void* QXmlQuery_NetworkAccessManager(void* ptr);
void* QXmlQuery_InitialTemplateName(void* ptr);
void* QXmlQuery_NamePool(void* ptr);
long long QXmlQuery_QueryLanguage(void* ptr);
char QXmlQuery_EvaluateTo2(void* ptr, void* callback);
char QXmlQuery_EvaluateTo4(void* ptr, void* target);
char QXmlQuery_EvaluateTo5(void* ptr, struct QtXmlPatterns_PackedString output);
char QXmlQuery_EvaluateTo3(void* ptr, struct QtXmlPatterns_PackedString target);
char QXmlQuery_IsValid(void* ptr);
void* QXmlQuery_UriResolver(void* ptr);
void QXmlQuery_EvaluateTo(void* ptr, void* result);
void* QXmlResultItems_Next(void* ptr);
void* QXmlResultItems_NewQXmlResultItems();
void QXmlResultItems_DestroyQXmlResultItems(void* ptr);
void QXmlResultItems_DestroyQXmlResultItemsDefault(void* ptr);
void* QXmlResultItems_Current(void* ptr);
char QXmlResultItems_HasError(void* ptr);
void* QXmlSchema_NewQXmlSchema();
void* QXmlSchema_NewQXmlSchema2(void* other);
char QXmlSchema_Load2(void* ptr, void* source, void* documentUri);
char QXmlSchema_Load3(void* ptr, void* data, void* documentUri);
char QXmlSchema_Load(void* ptr, void* source);
void QXmlSchema_SetMessageHandler(void* ptr, void* handler);
void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchema_SetUriResolver(void* ptr, void* resolver);
void QXmlSchema_DestroyQXmlSchema(void* ptr);
void* QXmlSchema_MessageHandler(void* ptr);
void* QXmlSchema_NetworkAccessManager(void* ptr);
void* QXmlSchema_DocumentUri(void* ptr);
void* QXmlSchema_NamePool(void* ptr);
char QXmlSchema_IsValid(void* ptr);
void* QXmlSchema_UriResolver(void* ptr);
void* QXmlSchemaValidator_NewQXmlSchemaValidator();
void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema);
void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler);
void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchemaValidator_SetSchema(void* ptr, void* schema);
void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver);
void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr);
void* QXmlSchemaValidator_MessageHandler(void* ptr);
void* QXmlSchemaValidator_NetworkAccessManager(void* ptr);
void* QXmlSchemaValidator_NamePool(void* ptr);
void* QXmlSchemaValidator_Schema(void* ptr);
char QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri);
char QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri);
char QXmlSchemaValidator_Validate(void* ptr, void* source);
void* QXmlSchemaValidator_UriResolver(void* ptr);
void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice);
void QXmlSerializer_AtomicValue(void* ptr, void* value);
void QXmlSerializer_AtomicValueDefault(void* ptr, void* value);
void QXmlSerializer_Attribute(void* ptr, void* name, void* value);
void QXmlSerializer_AttributeDefault(void* ptr, void* name, void* value);
void QXmlSerializer_Characters(void* ptr, void* value);
void QXmlSerializer_CharactersDefault(void* ptr, void* value);
void QXmlSerializer_Comment(void* ptr, struct QtXmlPatterns_PackedString value);
void QXmlSerializer_CommentDefault(void* ptr, struct QtXmlPatterns_PackedString value);
void QXmlSerializer_EndDocument(void* ptr);
void QXmlSerializer_EndDocumentDefault(void* ptr);
void QXmlSerializer_EndElement(void* ptr);
void QXmlSerializer_EndElementDefault(void* ptr);
void QXmlSerializer_EndOfSequence(void* ptr);
void QXmlSerializer_EndOfSequenceDefault(void* ptr);
void QXmlSerializer_NamespaceBinding(void* ptr, void* nb);
void QXmlSerializer_NamespaceBindingDefault(void* ptr, void* nb);
void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, struct QtXmlPatterns_PackedString value);
void QXmlSerializer_ProcessingInstructionDefault(void* ptr, void* name, struct QtXmlPatterns_PackedString value);
void QXmlSerializer_SetCodec(void* ptr, void* outputCodec);
void QXmlSerializer_StartDocument(void* ptr);
void QXmlSerializer_StartDocumentDefault(void* ptr);
void QXmlSerializer_StartElement(void* ptr, void* name);
void QXmlSerializer_StartElementDefault(void* ptr, void* name);
void QXmlSerializer_StartOfSequence(void* ptr);
void QXmlSerializer_StartOfSequenceDefault(void* ptr);
void* QXmlSerializer_OutputDevice(void* ptr);
void* QXmlSerializer_Codec(void* ptr);

#ifdef __cplusplus
}
#endif

#endif