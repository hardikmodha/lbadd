// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Error-1]
	_ = x[EOF-2]
	_ = x[StatementSeparator-3]
	_ = x[KeywordAbort-4]
	_ = x[KeywordAction-5]
	_ = x[KeywordAdd-6]
	_ = x[KeywordAfter-7]
	_ = x[KeywordAll-8]
	_ = x[KeywordAlter-9]
	_ = x[KeywordAnalyze-10]
	_ = x[KeywordAnd-11]
	_ = x[KeywordAs-12]
	_ = x[KeywordAsc-13]
	_ = x[KeywordAttach-14]
	_ = x[KeywordAutoincrement-15]
	_ = x[KeywordBefore-16]
	_ = x[KeywordBegin-17]
	_ = x[KeywordBetween-18]
	_ = x[KeywordBy-19]
	_ = x[KeywordCascade-20]
	_ = x[KeywordCase-21]
	_ = x[KeywordCast-22]
	_ = x[KeywordCheck-23]
	_ = x[KeywordCollate-24]
	_ = x[KeywordColumn-25]
	_ = x[KeywordCommit-26]
	_ = x[KeywordConflict-27]
	_ = x[KeywordConstraint-28]
	_ = x[KeywordCreate-29]
	_ = x[KeywordCross-30]
	_ = x[KeywordCurrent-31]
	_ = x[KeywordCurrentDate-32]
	_ = x[KeywordCurrentTime-33]
	_ = x[KeywordCurrentTimestamp-34]
	_ = x[KeywordDatabase-35]
	_ = x[KeywordDefault-36]
	_ = x[KeywordDeferrable-37]
	_ = x[KeywordDeferred-38]
	_ = x[KeywordDelete-39]
	_ = x[KeywordDesc-40]
	_ = x[KeywordDetach-41]
	_ = x[KeywordDistinct-42]
	_ = x[KeywordDo-43]
	_ = x[KeywordDrop-44]
	_ = x[KeywordEach-45]
	_ = x[KeywordElse-46]
	_ = x[KeywordEnd-47]
	_ = x[KeywordEscape-48]
	_ = x[KeywordExcept-49]
	_ = x[KeywordExclude-50]
	_ = x[KeywordExclusive-51]
	_ = x[KeywordExists-52]
	_ = x[KeywordExplain-53]
	_ = x[KeywordFail-54]
	_ = x[KeywordFilter-55]
	_ = x[KeywordFirst-56]
	_ = x[KeywordFollowing-57]
	_ = x[KeywordFor-58]
	_ = x[KeywordForeign-59]
	_ = x[KeywordFrom-60]
	_ = x[KeywordFull-61]
	_ = x[KeywordGlob-62]
	_ = x[KeywordGroup-63]
	_ = x[KeywordGroups-64]
	_ = x[KeywordHaving-65]
	_ = x[KeywordIf-66]
	_ = x[KeywordIgnore-67]
	_ = x[KeywordImmediate-68]
	_ = x[KeywordIn-69]
	_ = x[KeywordIndex-70]
	_ = x[KeywordIndexed-71]
	_ = x[KeywordInitially-72]
	_ = x[KeywordInner-73]
	_ = x[KeywordInsert-74]
	_ = x[KeywordInstead-75]
	_ = x[KeywordIntersect-76]
	_ = x[KeywordInto-77]
	_ = x[KeywordIs-78]
	_ = x[KeywordIsnull-79]
	_ = x[KeywordJoin-80]
	_ = x[KeywordKey-81]
	_ = x[KeywordLast-82]
	_ = x[KeywordLeft-83]
	_ = x[KeywordLike-84]
	_ = x[KeywordLimit-85]
	_ = x[KeywordMatch-86]
	_ = x[KeywordNatural-87]
	_ = x[KeywordNo-88]
	_ = x[KeywordNot-89]
	_ = x[KeywordNothing-90]
	_ = x[KeywordNotnull-91]
	_ = x[KeywordNull-92]
	_ = x[KeywordNulls-93]
	_ = x[KeywordOf-94]
	_ = x[KeywordOffset-95]
	_ = x[KeywordOn-96]
	_ = x[KeywordOr-97]
	_ = x[KeywordOrder-98]
	_ = x[KeywordOthers-99]
	_ = x[KeywordOuter-100]
	_ = x[KeywordOver-101]
	_ = x[KeywordPartition-102]
	_ = x[KeywordPlan-103]
	_ = x[KeywordPragma-104]
	_ = x[KeywordPreceding-105]
	_ = x[KeywordPrimary-106]
	_ = x[KeywordQuery-107]
	_ = x[KeywordRaise-108]
	_ = x[KeywordRange-109]
	_ = x[KeywordRecursive-110]
	_ = x[KeywordReferences-111]
	_ = x[KeywordRegexp-112]
	_ = x[KeywordReindex-113]
	_ = x[KeywordRelease-114]
	_ = x[KeywordRename-115]
	_ = x[KeywordReplace-116]
	_ = x[KeywordRestrict-117]
	_ = x[KeywordRight-118]
	_ = x[KeywordRollback-119]
	_ = x[KeywordRow-120]
	_ = x[KeywordRows-121]
	_ = x[KeywordSavepoint-122]
	_ = x[KeywordSelect-123]
	_ = x[KeywordSet-124]
	_ = x[KeywordTable-125]
	_ = x[KeywordTemp-126]
	_ = x[KeywordTemporary-127]
	_ = x[KeywordThen-128]
	_ = x[KeywordTies-129]
	_ = x[KeywordTo-130]
	_ = x[KeywordTransaction-131]
	_ = x[KeywordTrigger-132]
	_ = x[KeywordUnbounded-133]
	_ = x[KeywordUnion-134]
	_ = x[KeywordUnique-135]
	_ = x[KeywordUpdate-136]
	_ = x[KeywordUsing-137]
	_ = x[KeywordVacuum-138]
	_ = x[KeywordValues-139]
	_ = x[KeywordView-140]
	_ = x[KeywordVirtual-141]
	_ = x[KeywordWhen-142]
	_ = x[KeywordWhere-143]
	_ = x[KeywordWindow-144]
	_ = x[KeywordWith-145]
	_ = x[KeywordWithout-146]
	_ = x[Literal-147]
}

const _Type_name = "UnknownErrorEOFStatementSeparatorKeywordAbortKeywordActionKeywordAddKeywordAfterKeywordAllKeywordAlterKeywordAnalyzeKeywordAndKeywordAsKeywordAscKeywordAttachKeywordAutoincrementKeywordBeforeKeywordBeginKeywordBetweenKeywordByKeywordCascadeKeywordCaseKeywordCastKeywordCheckKeywordCollateKeywordColumnKeywordCommitKeywordConflictKeywordConstraintKeywordCreateKeywordCrossKeywordCurrentKeywordCurrentDateKeywordCurrentTimeKeywordCurrentTimestampKeywordDatabaseKeywordDefaultKeywordDeferrableKeywordDeferredKeywordDeleteKeywordDescKeywordDetachKeywordDistinctKeywordDoKeywordDropKeywordEachKeywordElseKeywordEndKeywordEscapeKeywordExceptKeywordExcludeKeywordExclusiveKeywordExistsKeywordExplainKeywordFailKeywordFilterKeywordFirstKeywordFollowingKeywordForKeywordForeignKeywordFromKeywordFullKeywordGlobKeywordGroupKeywordGroupsKeywordHavingKeywordIfKeywordIgnoreKeywordImmediateKeywordInKeywordIndexKeywordIndexedKeywordInitiallyKeywordInnerKeywordInsertKeywordInsteadKeywordIntersectKeywordIntoKeywordIsKeywordIsnullKeywordJoinKeywordKeyKeywordLastKeywordLeftKeywordLikeKeywordLimitKeywordMatchKeywordNaturalKeywordNoKeywordNotKeywordNothingKeywordNotnullKeywordNullKeywordNullsKeywordOfKeywordOffsetKeywordOnKeywordOrKeywordOrderKeywordOthersKeywordOuterKeywordOverKeywordPartitionKeywordPlanKeywordPragmaKeywordPrecedingKeywordPrimaryKeywordQueryKeywordRaiseKeywordRangeKeywordRecursiveKeywordReferencesKeywordRegexpKeywordReindexKeywordReleaseKeywordRenameKeywordReplaceKeywordRestrictKeywordRightKeywordRollbackKeywordRowKeywordRowsKeywordSavepointKeywordSelectKeywordSetKeywordTableKeywordTempKeywordTemporaryKeywordThenKeywordTiesKeywordToKeywordTransactionKeywordTriggerKeywordUnboundedKeywordUnionKeywordUniqueKeywordUpdateKeywordUsingKeywordVacuumKeywordValuesKeywordViewKeywordVirtualKeywordWhenKeywordWhereKeywordWindowKeywordWithKeywordWithoutLiteral"

var _Type_index = [...]uint16{0, 7, 12, 15, 33, 45, 58, 68, 80, 90, 102, 116, 126, 135, 145, 158, 178, 191, 203, 217, 226, 240, 251, 262, 274, 288, 301, 314, 329, 346, 359, 371, 385, 403, 421, 444, 459, 473, 490, 505, 518, 529, 542, 557, 566, 577, 588, 599, 609, 622, 635, 649, 665, 678, 692, 703, 716, 728, 744, 754, 768, 779, 790, 801, 813, 826, 839, 848, 861, 877, 886, 898, 912, 928, 940, 953, 967, 983, 994, 1003, 1016, 1027, 1037, 1048, 1059, 1070, 1082, 1094, 1108, 1117, 1127, 1141, 1155, 1166, 1178, 1187, 1200, 1209, 1218, 1230, 1243, 1255, 1266, 1282, 1293, 1306, 1322, 1336, 1348, 1360, 1372, 1388, 1405, 1418, 1432, 1446, 1459, 1473, 1488, 1500, 1515, 1525, 1536, 1552, 1565, 1575, 1587, 1598, 1614, 1625, 1636, 1645, 1663, 1677, 1693, 1705, 1718, 1731, 1743, 1756, 1769, 1780, 1794, 1805, 1817, 1830, 1841, 1855, 1862}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}