```
 checking cached actions
FATAL: bazel crashed due to an internal error. Printing stack trace:
java.lang.RuntimeException: Unrecoverable error while evaluating node 'SINGLE_EXTENSION_EVAL:ModuleExtensionId{bzlFileLabel=@@bazel_features~//private:extensions.bzl, extensionName=version_extension, isolationKey=Optional.empty}' (requested by nodes 'BZLMOD_REPO_RULE:@@bazel_features~~version_extension~bazel_features_globals', 'BZLMOD_REPO_RULE:@@bazel_features~~version_extension~bazel_features_version')
        at com.google.devtools.build.skyframe.AbstractParallelEvaluator$Evaluate.run(AbstractParallelEvaluator.java:550)
        at com.google.devtools.build.lib.concurrent.AbstractQueueVisitor$WrappedRunnable.run(AbstractQueueVisitor.java:414)
        at java.base/java.util.concurrent.ForkJoinTask$AdaptedRunnableAction.exec(Unknown Source)
        at java.base/java.util.concurrent.ForkJoinTask.doExec(Unknown Source)
        at java.base/java.util.concurrent.ForkJoinPool$WorkQueue.topLevelExec(Unknown Source)
        at java.base/java.util.concurrent.ForkJoinPool.scan(Unknown Source)
        at java.base/java.util.concurrent.ForkJoinPool.runWorker(Unknown Source)
        at java.base/java.util.concurrent.ForkJoinWorkerThread.run(Unknown Source)
Caused by: java.lang.NullPointerException: null value in entry: bazel_skylib=null
        at com.google.common.collect.CollectPreconditions.checkEntryNotNull(CollectPreconditions.java:33)
        at com.google.common.collect.ImmutableMapEntry.<init>(ImmutableMapEntry.java:54)
        at com.google.common.collect.ImmutableMap.entryOf(ImmutableMap.java:341)
        at com.google.common.collect.ImmutableMap$Builder.put(ImmutableMap.java:450)
        at com.google.devtools.build.lib.bazel.bzlmod.Module.getRepoMappingWithBazelDepsOnly(Module.java:67)
        at com.google.devtools.build.lib.bazel.bzlmod.BazelDepGraphFunction.getExtensionUsagesById(BazelDepGraphFunction.java:233)
        at com.google.devtools.build.lib.bazel.bzlmod.BazelDepGraphFunction.getExtensionUsagesById(BazelDepGraphFunction.java:221)
        at com.google.devtools.build.lib.bazel.bzlmod.SingleExtensionEvalFunction.tryGettingValueFromLockFile(SingleExtensionEvalFunction.java:258)
        at com.google.devtools.build.lib.bazel.bzlmod.SingleExtensionEvalFunction.compute(SingleExtensionEvalFunction.java:165)
        at com.google.devtools.build.skyframe.AbstractParallelEvaluator$Evaluate.run(AbstractParallelEvaluator.java:461)
        ... 7 more
```