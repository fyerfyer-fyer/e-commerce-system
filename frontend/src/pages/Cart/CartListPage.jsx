import React, { useState } from 'react';
import { Card, Table, Typography, Space, Button, InputNumber, Modal, message } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

const initialCartData = [
  {
    key: '1',
    product_id: 'P001',
    name: 'Smartphone',
    quantity: 2,
    price: 499.99,
  },
  {
    key: '2',
    product_id: 'P002',
    name: 'Laptop',
    quantity: 1,
    price: 999.99,
  },
];

const calculateTotalPrice = (items) => {
  return items.reduce((total, item) => total + item.quantity * item.price, 0).toFixed(2);
};

function CartListPage() {
  const [cartItems, setCartItems] = useState(initialCartData);
  const [totalPrice, setTotalPrice] = useState(calculateTotalPrice(initialCartData));

  // Handle quantity change for a specific item
  const handleQuantityChange = (item, newQuantity) => {
    if (newQuantity < 1) {
      message.error('Quantity must be at least 1.');
      return;
    }

    const updatedItems = cartItems.map((i) =>
      i.key === item.key ? { ...i, quantity: newQuantity } : i
    );

    setCartItems(updatedItems);
    setTotalPrice(calculateTotalPrice(updatedItems));
  };

  // Handle item deletion from the cart
  const handleDeleteItem = (item) => {
    Modal.confirm({
      title: 'Remove Item from Cart',
      content: `Are you sure you want to remove "${item.name}" from your cart?`,
      okText: 'Remove',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        const updatedItems = cartItems.filter((i) => i.key !== item.key);
        setCartItems(updatedItems);
        setTotalPrice(calculateTotalPrice(updatedItems));
        message.success('Item removed from cart!');
      },
    });
  };

  // Handle clearing the entire cart
  const handleClearCart = () => {
    Modal.confirm({
      title: 'Clear Cart',
      content: 'Are you sure you want to clear your entire cart?',
      okText: 'Clear',
      okType: 'danger',
      cancelText: 'Cancel',
      onOk() {
        setCartItems([]);
        setTotalPrice('0.00');
        message.success('Cart cleared!');
      },
    });
  };

  // Columns for the table
  const columns = [
    {
      title: 'Product ID',
      dataIndex: 'product_id',
      key: 'product_id',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
      render: (price) => `$${price.toFixed(2)}`,
    },
    {
      title: 'Quantity',
      dataIndex: 'quantity',
      key: 'quantity',
      render: (quantity, record) => (
        <InputNumber
          min={1}
          value={quantity}
          onChange={(value) => handleQuantityChange(record, value)}
          style={{ width: '100%' }}
        />
      ),
    },
    {
      title: 'Subtotal',
      key: 'subtotal',
      render: (_, record) => `$${(record.quantity * record.price).toFixed(2)}`,
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button type="link" danger onClick={() => handleDeleteItem(record)}>
            Remove
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="page-container">
      <Title level={2}>Shopping Cart</Title>
      <Card title="Your Cart" bordered={false}>
        {cartItems.length === 0 ? (
          <p>Your cart is empty.</p>
        ) : (
          <>
            <Table
              columns={columns}
              dataSource={cartItems}
              pagination={{ pageSize: 5 }}
              rowKey="key"
            />
            <div style={{ marginTop: 16, textAlign: 'right' }}>
              <strong>Total Price:</strong> ${totalPrice}
            </div>
            <Space style={{ marginTop: 16, display: 'flex', justifyContent: 'flex-end' }}>
              <Button
                type="primary"
                onClick={() => window.location.href = '/order/submit'}
              >
                Proceed to Checkout
              </Button>
              <Button type="danger" onClick={handleClearCart}>
                Clear Cart
              </Button>
            </Space>
          </>
        )}
      </Card>
    </div>
  );
}

export default CartListPage;
