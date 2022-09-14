// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint", DatabaseManagementDbManagementPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint_associated_database", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint_associated_databases", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoints", DatabaseManagementDbManagementPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_database_management_job_executions_status", DatabaseManagementJobExecutionsStatusDataSource())
	tfresource.RegisterDatasource("oci_database_management_job_executions_statuses", DatabaseManagementJobExecutionsStatusesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database", DatabaseManagementManagedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_addm_task", DatabaseManagementManagedDatabaseAddmTaskDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_addm_tasks", DatabaseManagementManagedDatabaseAddmTasksDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_alert_log_count", DatabaseManagementManagedDatabaseAlertLogCountDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_alert_log_counts", DatabaseManagementManagedDatabaseAlertLogCountsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_attention_log_count", DatabaseManagementManagedDatabaseAttentionLogCountDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_attention_log_counts", DatabaseManagementManagedDatabaseAttentionLogCountsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_group", DatabaseManagementManagedDatabaseGroupDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_groups", DatabaseManagementManagedDatabaseGroupsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_task", DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_finding", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_findings", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_set", DatabaseManagementManagedDatabaseSqlTuningSetDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_sets", DatabaseManagementManagedDatabaseSqlTuningSetsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user", DatabaseManagementManagedDatabaseUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_consumer_group_privilege", DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_consumer_group_privileges", DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_data_access_container", DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_data_access_containers", DatabaseManagementManagedDatabaseUserDataAccessContainersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_object_privilege", DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_object_privileges", DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_proxied_for_user", DatabaseManagementManagedDatabaseUserProxiedForUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_proxied_for_users", DatabaseManagementManagedDatabaseUserProxiedForUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_role", DatabaseManagementManagedDatabaseUserRoleDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_roles", DatabaseManagementManagedDatabaseUserRolesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_users", DatabaseManagementManagedDatabaseUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases", DatabaseManagementManagedDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_asm_properties", DatabaseManagementManagedDatabasesAsmPropertiesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_asm_property", DatabaseManagementManagedDatabasesAsmPropertyDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_database_parameter", DatabaseManagementManagedDatabasesDatabaseParameterDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_database_parameters", DatabaseManagementManagedDatabasesDatabaseParametersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_proxy_user", DatabaseManagementManagedDatabasesUserProxyUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_proxy_users", DatabaseManagementManagedDatabasesUserProxyUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_system_privilege", DatabaseManagementManagedDatabasesUserSystemPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_system_privileges", DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_execution", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_execution_script", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_executions", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_aggregations", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_operation", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_operations", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_table_statistics", DatabaseManagementManagedDatabaseTableStatisticsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_preferred_credential", DatabaseManagementManagedDatabasePreferredCredentialDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_preferred_credentials", DatabaseManagementManagedDatabasePreferredCredentialsDataSource())
}