// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

import "github.com/cockroachdb/cockroach/pkg/cmd/roachtest/registry"

// RegisterTests registers all tests to the Registry. This powers `roachtest run`.
func RegisterTests(r registry.Registry) {
	registerAWSDMS(r)
	registerAcceptance(r)
	registerActiveRecord(r)
	registerAllocator(r)
	registerAlterPK(r)
	registerAsyncpg(r)
	registerAutoUpgrade(r)
	registerBackup(r)
	registerBackupMixedVersion(r)
	registerBackupNodeShutdown(r)
	registerCDC(r)
	registerCDCMixedVersions(r)
	registerCancel(r)
	registerClearRange(r)
	registerClockJumpTests(r)
	registerClockMonotonicTests(r)
	registerConnectionLatencyTest(r)
	registerCopy(r)
	registerCopyFrom(r)
	registerCostFuzz(r)
	registerDecommission(r)
	registerDecommissionBench(r)
	registerDiskFull(r)
	registerDiskStalledDetection(r)
	registerDjango(r)
	registerDrain(r)
	registerDrop(r)
	registerEncryption(r)
	registerFixtures(r)
	registerFlowable(r)
	registerFollowerReads(r)
	registerGORM(r)
	registerGopg(r)
	registerGossip(r)
	registerHibernate(r, hibernateOpts)
	registerHibernate(r, hibernateSpatialOpts)
	registerHotSpotSplits(r)
	registerImportDecommissioned(r)
	registerImportMixedVersion(r)
	registerImportNodeShutdown(r)
	registerImportTPCC(r)
	registerImportTPCH(r)
	registerInconsistency(r)
	registerIndexes(r)
	registerJasyncSQL(r)
	registerJepsen(r)
	registerJobsMixedVersions(r)
	registerKV(r)
	registerKVBench(r)
	registerKVContention(r)
	registerKVGracefulDraining(r)
	registerKVMultiStoreWithOverload(r)
	registerKVQuiescenceDead(r)
	registerKVRangeLookups(r)
	registerKVScalability(r)
	registerKVSplits(r)
	registerKnex(r)
	registerLOQRecovery(r)
	registerLargeRange(r)
	registerLedger(r)
	registerLibPQ(r)
	registerLiquibase(r)
	registerLoadSplits(r)
	registerMultiTenantFairness(r)
	registerMultiTenantTPCH(r)
	registerMultiTenantUpgrade(r)
	registerNetwork(r)
	registerNodeJSPostgres(r)
	registerOverload(r)
	registerPebbleWriteThroughput(r)
	registerPebbleYCSB(r)
	registerPgjdbc(r)
	registerPgx(r)
	registerPop(r)
	registerPsycopg(r)
	registerQueue(r)
	registerQuitTransfersLeases(r)
	registerRebalanceLoad(r)
	registerReplicaGC(r)
	registerRestart(r)
	registerRestore(r)
	registerRestoreNodeShutdown(r)
	registerRoachmart(r)
	registerRoachtest(r)
	registerRubyPG(r)
	registerRustPostgres(r)
	registerSQLAlchemy(r)
	registerSQLSmith(r)
	registerSSTableCorruption(r)
	registerSchemaChangeBulkIngest(r)
	registerSchemaChangeDuringKV(r)
	registerSchemaChangeDuringTPCC1000(r)
	registerSchemaChangeIndexTPCC100(r)
	registerSchemaChangeIndexTPCC1000(r)
	registerSchemaChangeInvertedIndex(r)
	registerSchemaChangeMixedVersions(r)
	registerDeclSchemaChangeCompatMixedVersions(r)
	registerSchemaChangeRandomLoad(r)
	registerScrubAllChecksTPCC(r)
	registerScrubIndexOnlyTPCC(r)
	registerSecondaryIndexesMultiVersionCluster(r)
	registerSecure(r)
	registerSequelize(r)
	registerSlowDrain(r)
	registerSyncTest(r)
	registerSysbench(r)
	registerTLP(r)
	registerTPCC(r)
	registerTPCDSVec(r)
	registerTPCE(r)
	registerTPCHBench(r)
	registerTPCHConcurrency(r)
	registerTPCHVec(r)
	registerTypeORM(r)
	registerUnoptimizedQueryOracle(r)
	registerValidateSystemSchemaAfterVersionUpgrade(r)
	registerVersion(r)
	registerYCSB(r)
}

// RegisterBenchmarks registers all benchmarks to the registry. This powers `roachtest bench`.
//
// TODO(tbg): it's unclear that `roachtest bench` is that useful, perhaps we make everything
// a roachtest but use a `bench` tag to determine what tests to understand as benchmarks.
func RegisterBenchmarks(r registry.Registry) {
	registerIndexesBench(r)
	registerTPCCBench(r)
	registerKVBench(r)
	registerTPCHBench(r)
}
