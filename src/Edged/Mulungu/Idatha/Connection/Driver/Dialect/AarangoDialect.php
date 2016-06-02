<?php namespace Edged\Mulungu\Idatha\Connection\Driver\Dialect;

class AarangoDialect{
    const GET_VERSION = '/_db/%s/_api/version%s';
    const CREATE_DOCUMENT = '/_db/%s/_api/document?collection=%s';
    const FIND_DOCUMENT_BY_ID = '/_db/%s/_api/document/%s';
    const DELETE_DOCUMENT_BY_ID = '/_db/%s/_api/document/%s';
    const UPDATE_DOCUMENT_BY_ID = '/_db/%s/_api/document/%s';
    const PATCH_DOCUMENT_BY_ID = '/_db/%s/_api/document/%s';
    const CREATE_COLLECTION = '/_db/%s/_api/collection';
    const FIND_COLLECTION_BY_NAME = '/_db/%s/_api/collection/%s';
    const QUERY_SIMPLE_ALL = '/_db/%s/_api/simple/all';
    const QUERY_AQL = '/_db/%s/_api/cursor';
}