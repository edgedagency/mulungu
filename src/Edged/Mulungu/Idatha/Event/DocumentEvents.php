<?php namespace Edged\Mulungu\Idatha\Event;

class DocumentEvents {

    const BEFORE_SAVE = "idatha_document.before_save";
    const BEFORE_UPDATE = "idatha_document.before_update";
    const BEFORE_DELETE = "idatha_document.before_delete";
    const BEFORE_PATCH = "idatha_document.before_patch";

    const AFTER_SAVE = "idatha_document.after_save";
    const AFTER_UPDATE = "idatha_document.after_update";
    const AFTER_DELETE = "idatha_document.after_delete";
    const AFTER_PATCH = "idatha_document.after_patch";
}