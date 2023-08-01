// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig13) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig13) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig13) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig13) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig13) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig13) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig13) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig13) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig13) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig13) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig13) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig13) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig13) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig13) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig13) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig13) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig13) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig13) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig13) SetOldSnapshotThreshold(v *wrapperspb.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig13) SetWalLevel(v PostgresqlConfig13_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig13) SetSynchronousCommit(v PostgresqlConfig13_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig13) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig13) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig13) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig13) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig13) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig13) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig13) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig13) SetConstraintExclusion(v PostgresqlConfig13_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig13) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig13) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig13) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig13) SetForceParallelMode(v PostgresqlConfig13_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig13) SetClientMinMessages(v PostgresqlConfig13_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig13) SetLogMinMessages(v PostgresqlConfig13_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig13) SetLogMinErrorStatement(v PostgresqlConfig13_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig13) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig13) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig13) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig13) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig13) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig13) SetLogErrorVerbosity(v PostgresqlConfig13_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig13) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig13) SetLogStatement(v PostgresqlConfig13_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig13) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig13) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig13) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig13) SetDefaultTransactionIsolation(v PostgresqlConfig13_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig13) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig13) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig13) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig13) SetByteaOutput(v PostgresqlConfig13_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig13) SetXmlbinary(v PostgresqlConfig13_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig13) SetXmloption(v PostgresqlConfig13_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig13) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig13) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig13) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig13) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig13) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig13) SetBackslashQuote(v PostgresqlConfig13_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig13) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig13) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig13) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig13) SetOperatorPrecedenceWarning(v *wrapperspb.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig13) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig13) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig13) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig13) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig13) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig13) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig13) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig13) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig13) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig13) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig13) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig13) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig13) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig13) SetEnableBitmapscan(v *wrapperspb.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig13) SetEnableHashagg(v *wrapperspb.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig13) SetEnableHashjoin(v *wrapperspb.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig13) SetEnableIndexscan(v *wrapperspb.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig13) SetEnableIndexonlyscan(v *wrapperspb.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig13) SetEnableMaterial(v *wrapperspb.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig13) SetEnableMergejoin(v *wrapperspb.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig13) SetEnableNestloop(v *wrapperspb.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig13) SetEnableSeqscan(v *wrapperspb.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig13) SetEnableSort(v *wrapperspb.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig13) SetEnableTidscan(v *wrapperspb.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig13) SetMaxWorkerProcesses(v *wrapperspb.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig13) SetMaxParallelWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig13) SetMaxParallelWorkersPerGather(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig13) SetAutovacuumVacuumScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig13) SetAutovacuumAnalyzeScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig13) SetDefaultTransactionReadOnly(v *wrapperspb.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig13) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig13) SetEnableParallelAppend(v *wrapperspb.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig13) SetEnableParallelHash(v *wrapperspb.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig13) SetEnablePartitionPruning(v *wrapperspb.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig13) SetEnablePartitionwiseAggregate(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig13) SetEnablePartitionwiseJoin(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig13) SetJit(v *wrapperspb.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig13) SetMaxParallelMaintenanceWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig13) SetParallelLeaderParticipation(v *wrapperspb.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig13) SetVacuumCleanupIndexScaleFactor(v *wrapperspb.DoubleValue) {
	m.VacuumCleanupIndexScaleFactor = v
}

func (m *PostgresqlConfig13) SetLogTransactionSampleRate(v *wrapperspb.DoubleValue) {
	m.LogTransactionSampleRate = v
}

func (m *PostgresqlConfig13) SetPlanCacheMode(v PostgresqlConfig13_PlanCacheMode) {
	m.PlanCacheMode = v
}

func (m *PostgresqlConfig13) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig13) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig13) SetSharedPreloadLibraries(v []PostgresqlConfig13_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogMinDuration(v *wrapperspb.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogAnalyze(v *wrapperspb.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogBuffers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogTiming(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogTriggers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogVerbose(v *wrapperspb.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig13) SetAutoExplainLogNestedStatements(v *wrapperspb.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig13) SetAutoExplainSampleRate(v *wrapperspb.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig13) SetPgHintPlanEnableHint(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig13) SetPgHintPlanEnableHintTable(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig13) SetPgHintPlanDebugPrint(v PostgresqlConfig13_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig13) SetPgHintPlanMessageLevel(v PostgresqlConfig13_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfig13) SetHashMemMultiplier(v *wrapperspb.DoubleValue) {
	m.HashMemMultiplier = v
}

func (m *PostgresqlConfig13) SetLogicalDecodingWorkMem(v *wrapperspb.Int64Value) {
	m.LogicalDecodingWorkMem = v
}

func (m *PostgresqlConfig13) SetMaintenanceIoConcurrency(v *wrapperspb.Int64Value) {
	m.MaintenanceIoConcurrency = v
}

func (m *PostgresqlConfig13) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *PostgresqlConfig13) SetWalKeepSize(v *wrapperspb.Int64Value) {
	m.WalKeepSize = v
}

func (m *PostgresqlConfig13) SetEnableIncrementalSort(v *wrapperspb.BoolValue) {
	m.EnableIncrementalSort = v
}

func (m *PostgresqlConfig13) SetAutovacuumVacuumInsertThreshold(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumInsertThreshold = v
}

func (m *PostgresqlConfig13) SetAutovacuumVacuumInsertScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumInsertScaleFactor = v
}

func (m *PostgresqlConfig13) SetLogMinDurationSample(v *wrapperspb.Int64Value) {
	m.LogMinDurationSample = v
}

func (m *PostgresqlConfig13) SetLogStatementSampleRate(v *wrapperspb.DoubleValue) {
	m.LogStatementSampleRate = v
}

func (m *PostgresqlConfig13) SetLogParameterMaxLength(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLength = v
}

func (m *PostgresqlConfig13) SetLogParameterMaxLengthOnError(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLengthOnError = v
}

func (m *PostgresqlConfig13) SetPgQualstatsEnabled(v *wrapperspb.BoolValue) {
	m.PgQualstatsEnabled = v
}

func (m *PostgresqlConfig13) SetPgQualstatsTrackConstants(v *wrapperspb.BoolValue) {
	m.PgQualstatsTrackConstants = v
}

func (m *PostgresqlConfig13) SetPgQualstatsMax(v *wrapperspb.Int64Value) {
	m.PgQualstatsMax = v
}

func (m *PostgresqlConfig13) SetPgQualstatsResolveOids(v *wrapperspb.BoolValue) {
	m.PgQualstatsResolveOids = v
}

func (m *PostgresqlConfig13) SetPgQualstatsSampleRate(v *wrapperspb.DoubleValue) {
	m.PgQualstatsSampleRate = v
}

func (m *PostgresqlConfig13) SetMaxStackDepth(v *wrapperspb.Int64Value) {
	m.MaxStackDepth = v
}

func (m *PostgresqlConfig13) SetGeqo(v *wrapperspb.BoolValue) {
	m.Geqo = v
}

func (m *PostgresqlConfig13) SetGeqoThreshold(v *wrapperspb.Int64Value) {
	m.GeqoThreshold = v
}

func (m *PostgresqlConfig13) SetGeqoEffort(v *wrapperspb.Int64Value) {
	m.GeqoEffort = v
}

func (m *PostgresqlConfig13) SetGeqoPoolSize(v *wrapperspb.Int64Value) {
	m.GeqoPoolSize = v
}

func (m *PostgresqlConfig13) SetGeqoGenerations(v *wrapperspb.Int64Value) {
	m.GeqoGenerations = v
}

func (m *PostgresqlConfig13) SetGeqoSelectionBias(v *wrapperspb.DoubleValue) {
	m.GeqoSelectionBias = v
}

func (m *PostgresqlConfig13) SetGeqoSeed(v *wrapperspb.DoubleValue) {
	m.GeqoSeed = v
}

func (m *PostgresqlConfig13) SetPgTrgmSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmSimilarityThreshold = v
}

func (m *PostgresqlConfig13) SetPgTrgmWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmWordSimilarityThreshold = v
}

func (m *PostgresqlConfig13) SetPgTrgmStrictWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmStrictWordSimilarityThreshold = v
}

func (m *PostgresqlConfig13) SetMaxStandbyArchiveDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyArchiveDelay = v
}

func (m *PostgresqlConfig13) SetSessionDurationTimeout(v *wrapperspb.Int64Value) {
	m.SessionDurationTimeout = v
}

func (m *PostgresqlConfigSet13) SetEffectiveConfig(v *PostgresqlConfig13) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet13) SetUserConfig(v *PostgresqlConfig13) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet13) SetDefaultConfig(v *PostgresqlConfig13) {
	m.DefaultConfig = v
}
