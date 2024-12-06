

# 概念

HTML是用来描述网页的标记语言  

HTML标签：通常成对出现，前面的是开始标签，后面的是结束标签。  

HTML元素：包含开始标签与结束标签。  

只有body结构会在浏览器中显示。  

# 标题

h标签表示标题。h1表示一级标题，h2表示二级标题，以此类推。  

## 一级标签  

### 二级标签  

#### 三级标签  

##### 四级标签  

###### 五级标签  

###### 六级标签  

# 段落

段落标签：用p标签表示，表示一个段落。  

# 链接

链接标签：用a标签表示。

`<a href="1" target="2" rel="3" download="4" title="5"></a>`  
“1”里边写网址。

“2”表示打开方式。

“3”表示链接的关系。

“4”写下载文件名，可以提示下载并保存为指定文件名。

“5”写定义链接的额外信息，鼠标悬停时的提示信息  。

## target属性

### 1. **`_blank`**

- **含义**: 在新窗口或新标签页中打开链接。
- **使用场景**: 当你希望用户在点击链接时不离开当前页面，而是打开一个新的页面时使用，用户可以轻松返回到原始页面。

### 2. **`_self`**

- **含义**: 在当前窗口或标签页中打开链接（默认值）。
- **使用场景**: 当你希望用户在当前页面中导航到新链接时使用。大多数情况下，链接默认使用 `_self`，因此不需要显式指定。

### 3. **`_parent`**

- **含义**: 在父框架中打开链接。
- **使用场景**: 当页面嵌套在框架中时，使用 `_parent` 可以在父框架中加载链接，而不是在当前框架中。

### 4. **`_top`**

- **含义**: 在整个窗口中打开链接，取消任何框架。
- **使用场景**: 当你希望用户离开所有框架，返回到整个窗口的顶层时使用。这通常用于确保用户看到完整的页面，而不是在框架中查看内容。

### 总结

- **`_blank`**: 新窗口/标签页。
- **`_self`**: 当前窗口/标签页（默认）。
- **`_parent`**: 父框架。
- **`_top`**: 整个窗口，取消所有框架。

## rel属性

- `nofollow`: 搜索引擎不要追踪此链接，常用于外部链接  
- `noopener`: 防止新的浏览上下文访问window.opener属性和open方法。  
- `noreferrer`: 不要发送 referer header即不告诉目标网站你从哪来。  

# 图像

图片标签用img标签表示，src属性指定图片的路径。  

`<img src="可以直接放链接" alt="">`  

alt属性：当图片无法显示时，显示alt属性里的内容。还可以设置图片的宽度和高度。  

`<img src="图片路径" alt="无法显示" width="100" height="100">`  

# 换行

br标签可以换行。  

hr标签可以创建水平线。  

# 属性

HTML标签可以有属性，属性用来设置标签的一些特性。  

`<p id="describe"></p>`  

`<a href="">一个链接</a>`  

属性值应该用双引号或单引号括起来,如果属性值有双引号，则要用单引号括起来。  

# 注释

用`<!-- 注释 -->`可以注释  

注意：在HTML中多个空格或空行，在页面显示时会仅会显示一个空格  

# 文本样式

- `b`标签表示加粗  

  `<b>加粗</b>`

- `i`标签表示斜体  

  `<b>斜体</b>`

- `u`标签表示下划线  

  `<u>下划线</u>`

- `small`标签表示小号字体  

  `<small>小号字体</small>`

- `sub`，`sup`标签分别表示上标和下标 
  `<sub>上标</sub>`

  `<sup>下标</sup>`

# 表格

tr=table row 一行  

td=table data 一个单元格  

th=table header 表头  

在table后跟border起到定义边框的作用，border="1"表示边框宽度为1像素。  

read标签用于定义表格标题，tbody标签用于定义表格主体，tfoot标签用于定义表格脚注。  

```
    <table border="1">
        <thead> 
            <tr>
             <th>列标题1</th>
             <th>列标题2</th>
             <th>列标题3</th>
            </tr>
      </thead>
      <tbody>
        <tr>
            <td>元素11</td>
            <td>元素12</td>
            <td>元素13</td>
        </tr>
        <tr>
            <td>元素21</td>
            <td>元素22</td>
            <td>元素23</td>  
        </tr>
        <tr>
            <td>元素31</td>
            <td>元素32</td>
            <td>元素33</td>
        </tr>
      </tbody>
    </table>
```



# 列表

用`ul`标签表示无序列表，用`li`标签表示列表项。  

```
 <ul>
    <li>无序列表 1</li>
    <li>无序列表 2</li>
    <li>无序列表 3</li>
    <li>无序列表 4</li>
 </ul>
```



---

用`ol`标签表示有序列表，用`li`标签表示列表项。  

```
 <ol>
    <li>有序列表 1</li>
    <li>有序列表 2</li>
    <li>有序列表 3</li>
    <li>有序列表 4</li>
 </ol>
```



---

`dl`标签表示定义列表，`dt`标签表示定义标题，`dd`标签表示定义描述。  

```
 <dl>
    <dt>定义标题1</dt>
    <dd>自定义描述1</dd>
    <dd>自定义描述2</dd>
    <dd>自定义描述3</dd
 </dl>
```



---

# 列表

## 块元素

块级元素会从新行开始，并占据整行宽度，可以把包含其他块级元素和行内元素。  

常见的有div标签，p标签，h标签等。  

### div元素

div用于创建块级元素，便于组织页面结构与布局，没有语义，一般用于组织内容。  

```html
<div class="nav">
  <a href="#">链接1</a>
  <a href="#">链接2</a>
  <a href="#">链接3</a>
</div>
<div class="content">
  <h2>标题1</h2>
  <p>内容1</p>
</div>
```

## 行内元素

行内元素通常在同一行内呈现，不会独占一行。行内元素不能包含块级元素，但可以包含其他行内元素。  

常见的有a标签，span标签等。  

### span元素

span用于创建行内元素，可以对文本进行样式或标记。  

# 表单

form标签用于创建表单  

- action 元素定义了表单数据提交的目标 URL  
- method 元素定义了提交数据的 HTTP 方法  
- label 元素用于为表单元素添加标签  
- input标签用于输入  
  - type="text" 用于输入文本  
  - type="password" 用于输入密码  
  - type="radio" 用于创建单选按钮  
  - type="checkbox" 用于创建多选按钮  
  - type="submit" 用于提交表单  

```html
<form action="提交地址" method="post"> 
  <label for="username">用户名：</label>
  <input type="text" id="username" placeholder="请输入用户名"><br><br> 
  <label for="pwd">密码</label>
  <input type="password" id="pwd" placeholder="请输入密码"> <br><br>

  <label for="">性别</label><br><br>
  <input type="radio" name="gender">男<br><br>
  <input type="radio" name="gender">女<br><br>

  <label for="">爱好</label><br><br>
  <input type="checkbox" name="hobby">吃饭<br><br>
  <input type="checkbox" name="hobby">睡觉<br><br>
  <input type="checkbox" name="hobby">打豆豆<br><br>

  <input type="submit"><br><br>
</form>
```
