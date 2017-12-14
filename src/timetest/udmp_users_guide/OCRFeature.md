# OCR 支持

用户可以对已经保存在 UDMP 中的文件使用 OCR 服务，然后将返回结果保存在 **Document** 或者 **Page** 的 **Bizdata** 中。

需要确保有合适的 **OcrProcedure** 实现类。
用户可自行实现，或者使用 jar 中已经提供的实现类。
该实现类实现了 `parseResponse` 方法，用来从 OCR 返回结果中提取业务所需要的数据，并添加到 **Document** 的 **Bizdata** 中。

## 使用方法

用户 **App** 需要对目标 **Document** 有 **Write** 权限，并且有满足要求的 **OcrProcedure** 实现类。

1. 在配置文件中设置 OCR 服务访问地址（`ocr_scheme` 和 `ocr_host`）。
1. 初始化 OCR 配置。
1. 使用 **DocId** 获取 **Document** 对象，并用此对象实例化 **OcrProcedure** 对象。
1. 设定需要执行 OCR 服务的 **Page** 对象以及使用的 **OcrTemplate**。
1. 执行 `perform` 方法。

## 代码示例

具体说明参见 `OcrDemo.java`。

```java
OcrConfig.init(configPath);

Document doc = Document.find(docId);

OcrProcedure ocr = new Pab302Procedure(doc);
ocr.addTarget(1, OcrTemplate.ID_CARD)
   .addTarget(2, OcrTemplate.ID_CARD);
ocr.perform("name of operator");

for(Map.Entry<OcrProcedure.Target, OcrProcedure.Result> entry : ocr.getResults().entrySet()) {
    String msg = String.format("Page %d ---- %s%n", entry.getKey().pageNo, entry.getValue().state);
    System.out.println(msg);
}
```

**注意：**

1. 执行 `perform` 方法后，可以通过 `getResults` 方法获得每个文件的执行结果。
1. OCR 的原始结果会被记录在对应 **Page** 的 **Bizdata** 中。如果执行多次，该记录会被最新的一次结果所覆盖。
