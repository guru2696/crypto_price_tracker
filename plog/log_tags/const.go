package log_tags

var MESSAGE = &Tag{"message"}

var USER_ID = &Tag{"user_id"}
var CLIENT_ID = &Tag{"client_id"}

var RESPONSE_BODY = &Tag{"response_body"}
var REQUEST_BODY = &Tag{"request_body"}
var STATUS_CODE = &Tag{"status_code"}
var RETRY_COUNT = &Tag{"retry_count"}

var HEADER = &Tag{"header"}

var RECOVER = &Tag{"recover"}

var FILE_NAME = &Tag{"file_name"}
var FILE_PATH = &Tag{"file_path"}
var ROW_NUM = &Tag{"row_number"}
var COLUMN_NUM = &Tag{"column_number"}

var FLU_ID = &Tag{"flu_id"}
var JOB_ID = &Tag{"job_id"}
var FLU_IDs = &Tag{"flu_ids"}
var JOB_IDs = &Tag{"job_ids"}
var MASTER_FLU_ID = &Tag{"master_flu_id"}
var FLU_BUILD = &Tag{"flu_build"}
var FLU = &Tag{"flu"}
var PROJECT_ID = &Tag{"project_id"}
var WORKFLOW_ID = &Tag{"workflow_id"}
var PROJECT_TAG = &Tag{"project_tag"}
var STEP_ID = &Tag{"step_id"}
var STEP_TYPE = &Tag{"step_type"}
var BATCH_ID = &Tag{"batch_id"}
var BATCH_IDs = &Tag{"batch_ids"}
var SUBMISSION_BATCH_ID = &Tag{"submission_batch_id"}
var SUBMISSION_BATCH_SAMPLE_ASSOCIATOR_ID = &Tag{"submission_batch_sample_associator_id"}
var SUBMISSION_ID = &Tag{"submission_id"}
var QUESTION_ID = &Tag{"question_id"}

var POSTBACK_REQUEST = &Tag{"postback_request"}
var POSTBACK_RESPONSE = &Tag{"postback_reponse"}
var ERROR_CODE = &Tag{"error_code"}
var REFERENCE_ID = &Tag{"reference_id"}

//VECTOR CONSTANTS

var RELATIVE_PATH = &Tag{"relative_path"}
var REPO_ID = &Tag{"repo_id"}
var VECTOR_CLIENT_ID = &Tag{"vector_client_id"}
var VECTOR_FILE_ID = &Tag{"vector_file_id"}
var ANGEL_FLU_ID = &Tag{"angel_flu_id"}

var FLU_VIEW_ID = &Tag{"flu_view_id"}

var Ques_ID = &Tag{"q_id"}

// Client SVC
var JOB = &Tag{"job"}

// Alchemist
var AI_TASK_ID = &Tag{"ai_task_id"}
var REQUEST_ID = &Tag{"request_id"}
var AI_MODEL = &Tag{"ai_model"}

// Resource Sampling
var ResourceId = &Tag{"resource_id"}
var ResourceType = &Tag{"resource_type"}
var AssociationId = &Tag{"association_id"}
