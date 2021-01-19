// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *PostgresqlConfig11) SetMaxConnections(v *wrappers.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig11) SetSharedBuffers(v *wrappers.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig11) SetTempBuffers(v *wrappers.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig11) SetMaxPreparedTransactions(v *wrappers.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig11) SetWorkMem(v *wrappers.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig11) SetMaintenanceWorkMem(v *wrappers.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig11) SetAutovacuumWorkMem(v *wrappers.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig11) SetTempFileLimit(v *wrappers.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig11) SetVacuumCostDelay(v *wrappers.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageHit(v *wrappers.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageMiss(v *wrappers.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageDirty(v *wrappers.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig11) SetVacuumCostLimit(v *wrappers.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig11) SetBgwriterDelay(v *wrappers.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig11) SetBgwriterLruMaxpages(v *wrappers.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig11) SetBgwriterLruMultiplier(v *wrappers.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig11) SetBgwriterFlushAfter(v *wrappers.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig11) SetBackendFlushAfter(v *wrappers.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig11) SetOldSnapshotThreshold(v *wrappers.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig11) SetWalLevel(v PostgresqlConfig11_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig11) SetSynchronousCommit(v PostgresqlConfig11_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig11) SetCheckpointTimeout(v *wrappers.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig11) SetCheckpointCompletionTarget(v *wrappers.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig11) SetCheckpointFlushAfter(v *wrappers.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig11) SetMaxWalSize(v *wrappers.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig11) SetMinWalSize(v *wrappers.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig11) SetMaxStandbyStreamingDelay(v *wrappers.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig11) SetDefaultStatisticsTarget(v *wrappers.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig11) SetConstraintExclusion(v PostgresqlConfig11_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig11) SetCursorTupleFraction(v *wrappers.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig11) SetFromCollapseLimit(v *wrappers.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig11) SetJoinCollapseLimit(v *wrappers.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig11) SetForceParallelMode(v PostgresqlConfig11_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig11) SetClientMinMessages(v PostgresqlConfig11_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig11) SetLogMinMessages(v PostgresqlConfig11_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig11) SetLogMinErrorStatement(v PostgresqlConfig11_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig11) SetLogMinDurationStatement(v *wrappers.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig11) SetLogCheckpoints(v *wrappers.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig11) SetLogConnections(v *wrappers.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig11) SetLogDisconnections(v *wrappers.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig11) SetLogDuration(v *wrappers.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig11) SetLogErrorVerbosity(v PostgresqlConfig11_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig11) SetLogLockWaits(v *wrappers.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig11) SetLogStatement(v PostgresqlConfig11_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig11) SetLogTempFiles(v *wrappers.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig11) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig11) SetRowSecurity(v *wrappers.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig11) SetDefaultTransactionIsolation(v PostgresqlConfig11_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig11) SetStatementTimeout(v *wrappers.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig11) SetLockTimeout(v *wrappers.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig11) SetIdleInTransactionSessionTimeout(v *wrappers.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig11) SetByteaOutput(v PostgresqlConfig11_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig11) SetXmlbinary(v PostgresqlConfig11_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig11) SetXmloption(v PostgresqlConfig11_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig11) SetGinPendingListLimit(v *wrappers.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig11) SetDeadlockTimeout(v *wrappers.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig11) SetMaxLocksPerTransaction(v *wrappers.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig11) SetMaxPredLocksPerTransaction(v *wrappers.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig11) SetArrayNulls(v *wrappers.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig11) SetBackslashQuote(v PostgresqlConfig11_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig11) SetDefaultWithOids(v *wrappers.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig11) SetEscapeStringWarning(v *wrappers.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig11) SetLoCompatPrivileges(v *wrappers.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig11) SetOperatorPrecedenceWarning(v *wrappers.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig11) SetQuoteAllIdentifiers(v *wrappers.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig11) SetStandardConformingStrings(v *wrappers.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig11) SetSynchronizeSeqscans(v *wrappers.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig11) SetTransformNullEquals(v *wrappers.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig11) SetExitOnError(v *wrappers.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig11) SetSeqPageCost(v *wrappers.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig11) SetRandomPageCost(v *wrappers.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig11) SetAutovacuumMaxWorkers(v *wrappers.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumCostDelay(v *wrappers.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumCostLimit(v *wrappers.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig11) SetAutovacuumNaptime(v *wrappers.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig11) SetArchiveTimeout(v *wrappers.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig11) SetTrackActivityQuerySize(v *wrappers.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig11) SetEnableBitmapscan(v *wrappers.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig11) SetEnableHashagg(v *wrappers.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig11) SetEnableHashjoin(v *wrappers.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig11) SetEnableIndexscan(v *wrappers.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig11) SetEnableIndexonlyscan(v *wrappers.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig11) SetEnableMaterial(v *wrappers.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig11) SetEnableMergejoin(v *wrappers.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig11) SetEnableNestloop(v *wrappers.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig11) SetEnableSeqscan(v *wrappers.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig11) SetEnableSort(v *wrappers.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig11) SetEnableTidscan(v *wrappers.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig11) SetMaxWorkerProcesses(v *wrappers.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig11) SetMaxParallelWorkers(v *wrappers.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig11) SetMaxParallelWorkersPerGather(v *wrappers.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumScaleFactor(v *wrappers.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig11) SetAutovacuumAnalyzeScaleFactor(v *wrappers.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig11) SetDefaultTransactionReadOnly(v *wrappers.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig11) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig11) SetEnableParallelAppend(v *wrappers.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig11) SetEnableParallelHash(v *wrappers.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig11) SetEnablePartitionPruning(v *wrappers.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig11) SetEnablePartitionwiseAggregate(v *wrappers.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig11) SetEnablePartitionwiseJoin(v *wrappers.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig11) SetJit(v *wrappers.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig11) SetMaxParallelMaintenanceWorkers(v *wrappers.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig11) SetParallelLeaderParticipation(v *wrappers.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig11) SetVacuumCleanupIndexScaleFactor(v *wrappers.DoubleValue) {
	m.VacuumCleanupIndexScaleFactor = v
}

func (m *PostgresqlConfig11) SetEffectiveIoConcurrency(v *wrappers.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig11) SetEffectiveCacheSize(v *wrappers.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig11) SetSharedPreloadLibraries(v []PostgresqlConfig11_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogMinDuration(v *wrappers.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogAnalyze(v *wrappers.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogBuffers(v *wrappers.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogTiming(v *wrappers.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogTriggers(v *wrappers.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogVerbose(v *wrappers.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogNestedStatements(v *wrappers.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig11) SetAutoExplainSampleRate(v *wrappers.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig11) SetPgHintPlanEnableHint(v *wrappers.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig11) SetPgHintPlanEnableHintTable(v *wrappers.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig11) SetPgHintPlanDebugPrint(v PostgresqlConfig11_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig11) SetPgHintPlanMessageLevel(v PostgresqlConfig11_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfigSet11) SetEffectiveConfig(v *PostgresqlConfig11) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet11) SetUserConfig(v *PostgresqlConfig11) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet11) SetDefaultConfig(v *PostgresqlConfig11) {
	m.DefaultConfig = v
}
