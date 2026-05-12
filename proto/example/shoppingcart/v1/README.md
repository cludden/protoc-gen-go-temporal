

<a name="example-shoppingcart-v1"></a>
# example.shoppingcart.v1

<a name="example-shoppingcart-v1-services"></a>
## Services

<a name="example-shoppingcart-v1-shoppingcart"></a>
## example.shoppingcart.v1.ShoppingCart

<a name="example-shoppingcart-v1-shoppingcart-workflows"></a>
### Workflows

---
<a name="example-shoppingcart-v1-shoppingcart-workflow"></a>
### example.shoppingcart.v1.ShoppingCart

**Input:** [example.shoppingcart.v1.ShoppingCartInput](#example-shoppingcart-v1-shoppingcartinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>

**Output:** [example.shoppingcart.v1.ShoppingCartOutput](#example-shoppingcart-v1-shoppingcartoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>example.shoppingcart.v1.ShoppingCart/${! nanoid() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>

**Queries:**

<table>
<tr><th>Query</th></tr>
<tr><td><a href="#example-shoppingcart-v1-shoppingcart-describe-query">example.shoppingcart.v1.ShoppingCart.Describe</a></td></tr>
</table>

**Signals:**

<table>
<tr><th>Signal</th><th>Start</th></tr>
<tr><td><a href="#example-shoppingcart-v1-shoppingcart-checkout-signal">example.shoppingcart.v1.ShoppingCart.Checkout</a></td><td>false</td></tr>
</table>

**Updates:**

<table>
<tr><th>Update</th></tr>
<tr><td><a href="#example-shoppingcart-v1-shoppingcart-updatecart-update">example.shoppingcart.v1.ShoppingCart.UpdateCart</a></td></tr>
</table>  

<a name="example-shoppingcart-v1-shoppingcart-queries"></a>
### Queries

---
<a name="example-shoppingcart-v1-shoppingcart-describe-query"></a>
### example.shoppingcart.v1.ShoppingCart.Describe



**Input:** [example.shoppingcart.v1.DescribeInput](#example-shoppingcart-v1-describeinput)



**Output:** [example.shoppingcart.v1.DescribeOutput](#example-shoppingcart-v1-describeoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>  

<a name="example-shoppingcart-v1-shoppingcart-signals"></a>
### Signals

---
<a name="example-shoppingcart-v1-shoppingcart-checkout-signal"></a>
### example.shoppingcart.v1.ShoppingCart.Checkout



**Input:** [example.shoppingcart.v1.CheckoutInput](#example-shoppingcart-v1-checkoutinput)

  

<a name="example-shoppingcart-v1-shoppingcart-updates"></a>
### Updates

---
<a name="example-shoppingcart-v1-shoppingcart-updatecart-update"></a>
### example.shoppingcart.v1.ShoppingCart.UpdateCart



**Input:** [example.shoppingcart.v1.UpdateCartInput](#example-shoppingcart-v1-updatecartinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>action</td>
<td><a href="#example-shoppingcart-v1-updatecartaction">example.shoppingcart.v1.UpdateCartAction</a></td>
<td><pre>
json_name: action
go_name: Action</pre></td>
</tr><tr>
<td>item_id</td>
<td>string</td>
<td><pre>
json_name: itemId
go_name: ItemId</pre></td>
</tr>
</table>

**Output:** [example.shoppingcart.v1.UpdateCartOutput](#example-shoppingcart-v1-updatecartoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>validate</td><td>true</td></tr>
</table> 

<a name="example-shoppingcart-v1-messages"></a>
## Messages

<a name="example-shoppingcart-v1-cartstate"></a>
### example.shoppingcart.v1.CartState

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>items</td>
<td><a href="#example-shoppingcart-v1-cartstate-itemsentry">example.shoppingcart.v1.CartState.ItemsEntry</a></td>
<td><pre>
json_name: items
go_name: Items</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-cartstate-itemsentry"></a>
### example.shoppingcart.v1.CartState.ItemsEntry

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>key</td>
<td>string</td>
<td><pre>
json_name: key
go_name: Key</pre></td>
</tr><tr>
<td>value</td>
<td>int32</td>
<td><pre>
json_name: value
go_name: Value</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-checkoutinput"></a>
### example.shoppingcart.v1.CheckoutInput



<a name="example-shoppingcart-v1-describeinput"></a>
### example.shoppingcart.v1.DescribeInput



<a name="example-shoppingcart-v1-describeoutput"></a>
### example.shoppingcart.v1.DescribeOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-shoppingcartinput"></a>
### example.shoppingcart.v1.ShoppingCartInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-shoppingcartoutput"></a>
### example.shoppingcart.v1.ShoppingCartOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-updatecartaction"></a>
### example.shoppingcart.v1.UpdateCartAction

<table>
<tr><th>Value</th><th>Description</th></tr>
<tr>
<td>UPDATE_CART_ACTION_UNSPECIFIED</td>
<td></td>
</tr><tr>
<td>UPDATE_CART_ACTION_ADD</td>
<td></td>
</tr><tr>
<td>UPDATE_CART_ACTION_REMOVE</td>
<td></td>
</tr>
</table>

<a name="example-shoppingcart-v1-updatecartinput"></a>
### example.shoppingcart.v1.UpdateCartInput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>action</td>
<td><a href="#example-shoppingcart-v1-updatecartaction">example.shoppingcart.v1.UpdateCartAction</a></td>
<td><pre>
json_name: action
go_name: Action</pre></td>
</tr><tr>
<td>item_id</td>
<td>string</td>
<td><pre>
json_name: itemId
go_name: ItemId</pre></td>
</tr>
</table>



<a name="example-shoppingcart-v1-updatecartoutput"></a>
### example.shoppingcart.v1.UpdateCartOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart</td>
<td><a href="#example-shoppingcart-v1-cartstate">example.shoppingcart.v1.CartState</a></td>
<td><pre>
json_name: cart
go_name: Cart</pre></td>
</tr>
</table>

