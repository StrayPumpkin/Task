

# CSS 格式

CSS用于定义网页样式与布局样式

```
选择器{
属性1:属性值1;  

属性2:属性值2;  

}
```

## CSS三种导入方式

- **内联样式**：将css样式直接放在到HTML标签中。
- **内部样式表**：直接在head标签中定义中定义。
- **外部样式表**：将css样式单独定义成一个文件，然后在head标签中通过link标签引入到HTML文档中。

**CSS样式优先级：内联样式 > 内部样式表 > 外部样式表**

## CSS选择器

CSS选择器是用来选择HTML元素的，通过选择器可以对HTML元素进行样式设置。

常用选择器：元素选择器、类选择器、ID选择器、通用选择器、子元素选择器、后代选择器、并集选择器、交集选择器、伪类选择器。

**它们的优先级为：行内样式 > ID选择器 > 类选择器 > 元素选择器。**

- **元素选择器**

元素选择器可以选择HTML文档中的具体元素，语法为“标签名”，选择页面上所有这种类型的标签。

+ **类选择器**

类选择器可以选择HTML文档中具有相同类属性的元素，语法为“.类名”，灵活！

+ **ID选择器**

ID选择器可以选择HTML文档中具有相同ID属性的元素，语法为“#ID名”。

+ **通用选择器**

通用选择器可以选择HTML文档中的所有元素，语法为“*”，一般用于样式初始化。

+ **子元素选择器**

子元素选择器可以选择HTML文档中某个元素的直接子元素，语法为“标签名 > 子元素标签名”。

+ **后代选择器**

后代选择器可以选择HTML文档中某个元素的所有后代元素，语法为“标签名 标签名”。

+ **并集选择器**

并集选择器可以选择HTML文档中多个选择器的元素，语法为“标签名1,标签名2”。

+ **交集选择器**

交集选择器可以选择HTML文档中多个选择器的交集元素，语法为“标签名1 标签名2”。

+ **相邻兄弟选择器**

相邻兄弟选择器可以选择HTML文档中某个元素的相邻的兄弟元素，语法为“标签名 + 标签名”，只能向下选择。

+ **通用兄弟选择器**

通用兄弟选择器可以选择HTML文档中某个元素的兄弟元素，语法为“标签名 ~ 标签名”。



## 字体属性

+ **color**

`color：英文表述/#十六进制/rgb(内填三个数，用逗号隔开)/rgba(内填4个数，最后一个数表示透明度范围在0~1)'`

+ **font-weight文本粗细**

`font-weight：bold/bolder/lighter/100-900，400为默认，700为加粗。`

+ **font-style文本样式**

`font-style：normal/italic。`

+ **font-family字体系列**

`font-family：英文表述/字体名称。`

## 背景属性

+ **background-image背景图片**

`background-image：url(图片路径)。`

+ **background-repeat背景平铺**

`background-repeat：repeat/repeat-x/repeat-y/no-repeat。`

+ **background-position背景位置**

`background-position：x轴偏移量/y轴偏移量。`

+ **background-size背景尺寸**

`background-size：数值/百分比/cover 完全覆盖整个容器，会有切割/contain 保持比例缩放，无切割。`

+ **background-position背景位置**

`background-position：左中右 上中下/百分比 百分比。`

## 文本属性

+ **text-align文本对齐方式**

`text-align：left/right/center/justify。`

+ **text-decoration文本装饰**

`text-decoration：none/underline/overline/line-through/blink。`

+ **text-indent文本缩进**

`text-indent：数值/百分比。`

+ **text-transform文本转换**

`text-transform：none/capitalize 每个单词开头大写/uppercase 全部大写/lowercase 全部小写。`

## 表格属性

+ **border 边框**

`border：边框宽度/边框样式/边框颜色。`

+ **border-collapse折叠边框**

`border-collapse：collapse/separate。`

+ **表格文字对齐**

在对td标签设置水平对齐用text-align，垂直对齐用vertical-align。

+ **padding表格填充**

可以不用设置宽高，只设置padding。

## 盒子模型

+ **width宽度**

`width：数值/百分比。`

+ **height高度**

`height：数值/百分比。`

+ **padding内边距**

`padding：数值/百分比，上下一个值，左右一个值。也可分别设置各个方向的值。撑大内容区，使内容更好展示。`

+ **border边框**

`border：边框宽度/边框样式/边框颜色。`

+ **margin外边距**

`margin：数值/百分比，是透明的，用来调整元素之间的距离。`

## 弹性盒模型

针对大盒子里的小盒子。在大盒子里设置`display: flex;`，默认是水平排列。

### 父元素属性

+ **flex-direction子元素排列方向**

`flex-direction：row横向排列/row-reverse反向横向排列/column纵向排列/column-reverse反向纵向排列。`

+ **justify-content垂直方向**

`justify-content：flex-start/flex-end/center/space-between/space-around。`

+ **align-items水平方向**

`align-items：flex-start/flex-end/center/baseline/stretch。`

### 子元素属性

+ ***flex-grow/flex放大比例***

`flex-grow：数值，根据权重分配剩余空间。设置了flex后宽度失效。`

## 浮动

浮动可以使元素脱离文档流，并向左或向右移动，不能向上或向下移动直到遇到另一个浮动元素或碰到父元素的边框。  

`float：left/right/none。`

## 清除浮动

清除浮动可以解决浮动元素的高度塌陷问题。

- 方法一：父元素设置高度。
- 方法二：父元素设置`overflow：hidden`。
- 方法三：受影响元素设置`clear：both`。

## 定位

定位可以将元素从文档流中移除，并相对于其包含块进行定位。  

`position：static/relative/absolute/fixed`

+ ***relative相对定位***

相对定位可以相对于元素自身进行定位，没有脱离文档流。  

`top/bottom/left/right：数值/百分比`

+ ***absolute绝对定位***

绝对定位可以相对于距离最近的已定位祖先元素进行定位，脱离文档流，每设置一个定位属性，就多一层。  `top/bottom/left/right：数值/百分比`

+ ***fixed固定定位***

固定定位可以相对于浏览器窗口进行定位，脱离文档流。  

`top/bottom/left/right：数值/百分比`

### z-index层级

`z-index：数值，越大越在前面。`

## 新特性

+ **圆角**

`border-radius：数值/百分比，可以设置四个角的圆角。左上 右下100%表示圆形`

+ **阴影**

`box-shadow：h-shadow水平偏移量/v-shadow垂直偏移量/blur模糊距离/color阴影颜色，设置阴影效果`

+ **动画**

在style中创建@keyframes动画名称。

```
@keyframes 动画名称 {
    from/0% {
        css样式
    }
    percent {
        css样式
    }
    to/100% {
        css样式
    }
}
```

```
```animation属性，设置name动画名称，duration动画时间，timing-function动画速度曲线，delay动画延迟，iteration-count动画次数，direction动画方向。 ``` 
```

+ animation-play-state用于设置动画播放状态，paused暂停，running播放。  
+ time-function的选择：ease逐渐变慢/linear匀速/ease-in加速/ease-out减速/ease-in-out先加速后减速
+ direction的选择：normal正常/reverse反向/alternate交替。  
+ hover，鼠标划上去，动画效果。  
+ opacity表示透明度，0~1，1表示完全不透明。

### 媒体查询

可以针对不同的屏幕尺寸设置不同的样式。  在head标签里添加meta标签。  

媒体类型：all/screen/print。  

媒体特征：min-width/max-width/min-height/max-height/orientation。  

例子：

```
@media screen and (max-width: 600px) {
    /* 针对屏幕宽度小于600px的样式 */
}
```

### 字体图标

推荐：[阿里字体图标库](https://www.iconfont.cn/)  

选择下载代码。  

推荐用 font class 方式使用字体图标，可以自定义颜色和大小。

1. 引入生成的fontclass代码

   ```html
   <link rel="stylesheet" href="https://at.alicdn.com/t/font_1674484_67y171191.css">
   ```

2. 挑选相应的图标并获取类名，应用于页面。

   ```html
   <span class="iconfont icon-xxx"></span>
   ```

