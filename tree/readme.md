## 树

树（tree）是包含n（n>0）个结点的有穷集，其中：
（1）每个元素称为结点（node）；
（2）有一个特定的结点被称为根结点或树根（root）。
（3）除根结点之外的其余数据元素被分为m（m≥0）个互不相交的集合T1，T2，……Tm-1，其中每一个集合Ti（1<=i<=m）本身也是一棵树，被称作原树的子树（subtree）。

![树](https://upload.wikimedia.org/wikipedia/commons/7/7e/Treedatastructure.png)

### 相关术语：

* 节点的度：一个节点含有的子树的个数称为该节点的度；
* 树的度：一棵树中，最大的节点的度称为树的度；
* 叶节点或终端节点：度为零的节点；
* 非终端节点或分支节点：度不为零的节点；
* 父亲节点或父节点：若一个节点含有子节点，则这个节点称为其子节点的父节点；
* 孩子节点或子节点：一个节点含有的子树的根节点称为该节点的子节点；
* 兄弟节点：具有相同父节点的节点互称为兄弟节点；
* 节点的层次：从根开始定义起，根为第1层，根的子节点为第2层，以此类推；
* 深度：对于任意节点n,n的深度为从根到n的唯一路径长，根的深度为0；
* 高度：对于任意节点n,n的高度为从n到一片树叶的最长路径长，所有树叶的高度为0；
* 堂兄弟节点：父节点在同一层的节点互为堂兄弟；
* 节点的祖先：从根到该节点所经分支上的所有节点；
* 子孙：以某节点为根的子树中任一节点都称为该节点的子孙。
* 森林：由m（m>=0）棵互不相交的树的集合称为森林；

### 二叉查找树(BST)

    若左子树不空，则左子树上所有结点的值均小于它的根结点的值；
    若右子树不空，则右子树上所有结点的值均大于或等于它的根结点的值；
    左、右子树也分别为二叉排序树；
    没有键值相等的节点;


### 平衡二叉树(AVL)

平衡二叉树（balanced binary tree）,又称 AVL 树。它或者是一棵空树,或者是具有如下性质的二叉树：

    它的左子树和右子树都是平衡二叉树，
    左子树和右子树的深度之差的绝对值不超过1。

平衡二叉树是对二叉搜索树(又称为二叉排序树)的一种改进。二叉搜索树有一个缺点就是，树的结构是无法预料的，随意性很大，它只与节点的值和插入的顺序有关系，往往得到的是一个不平衡的二叉树。在最坏的情况下，可能得到的是一个单支二叉树，其高度和节点数相同，相当于一个单链表，对其正常的时间复杂度有O(log(n))变成了O(n)，从而丧失了二叉排序树的一些应该有的优点。

### B树

BTree是平衡搜索 ***多叉树*** ，设树的度为2d（d>1），高度为h，那么BTree要满足以下条件：

    每个叶子结点的高度一样，等于h；
    每个非叶子结点由n-1个key和n个指针point组成，其中d<=n<=2d,key和point相互间隔，结点两端一定是key；
    叶子结点指针都为null；
    非叶子结点的key都是[key,data]二元组，其中key表示作为索引的键，data为键值所在行的数据；

### B+树

B+Tree是BTree的一个变种，设d为树的度数，h为树的高度，B+Tree和BTree的不同主要在于：

    B+Tree中的非叶子结点不存储数据，只存储键值；
    B+Tree的叶子结点没有指针，所有键值都会出现在叶子结点上，且key存储的键值对应data数据的物理地址；
    B+Tree的每个非叶子节点由n个键值key和n个指针point组成；


<img src="https://upload-images.jianshu.io/upload_images/3739895-f763877cf21aa94d.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/569/format/webp"/>

### B*树

B*树是B+树的变体，在B+树的非根和非叶子结点再增加指向兄弟的指针：

<img src="https://upload-images.jianshu.io/upload_images/3739895-d51e5dc039ffb536.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/569/format/webp"/>

### 其他

    二叉树 => 满二叉树，完全二叉树，哈夫曼树