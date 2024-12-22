import React, { useState, useEffect } from 'react';
import { Form, Input, Button, Typography, Select, Space, Card, Table, message } from 'antd';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../utils/auth';

const { Title } = Typography;
const { Option } = Select;

// Simulated cart data (this will be replaced with actual data from the backend)
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

// Simulated user addresses (this will be replaced with actual data from the backend)
const initialAddresses = [
  {
    id: '1',
    address_line1: '123 Main St',
    city: 'New York',
    state: 'NY',
    postal_code: '10001',
    country: 'USA',
  },
  {
    id: '2',
    address_line1: '456 Elm St',
    city: 'Los Angeles',
    state: 'CA',
    postal_code: '90001',
    country: 'USA',
  },
];

function OrderSubmitPage() {
  const [cartItems, setCartItems] = useState(initialCartData);
  const [addresses, setAddresses] = useState(initialAddresses);
  const [totalPrice, setTotalPrice] = useState(calculateTotalPrice(initialCartData));
  const [selectedAddress, setSelectedAddress] = useState(undefined);
  const [form] = Form.useForm();
  const { user } = useAuth();
  const navigate = useNavigate();

  // Function to calculate the total price of all items in the cart
  const calculateTotalPrice = (items) => {
    return items.reduce((total, item) => total + item.quantity * item.price, 0).toFixed(2);
  };

  // Handle form submission
  const onFinish = async (values) => {
    if (!selectedAddress) {
      message.error('Please select a delivery address.');
      return;
    }

    console.log('Submit order form values:', values);

    // Simulate a submit order request to the backend ( Placeholder for now )
    try {
      // Here you can add the logic to submit the order to the backend
      message.success('Order submitted successfully!');
      navigate('/order/list'); // Redirect to order list page after successful submission
    } catch (error) {
      message.error('Failed to submit order. Please try again.');
      console.error('Submit order error:', error);
    }
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
      title: 'Quantity',
      dataIndex: 'quantity',
      key: 'quantity',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
      render: (text) => `$${text.toFixed(2)}`,
    },
    {
      title: 'Subtotal',
      key: 'subtotal',
      render: (_, record) => `$${(record.quantity * record.price).toFixed(2)}`,
    },
  ];

  // Fetch user addresses from the backend (Placeholder for now)
  useEffect(() => {
    // Simulate fetching user addresses from the backend
    setTimeout(() => {
      setAddresses(initialAddresses);
    }, 1000);
  }, []);

  return (
    <div className="page-container">
      <Title level={2}>Submit Order</Title>
      <Card title="Order Summary" bordered={false}>
        {cartItems.length === 0 ? (
          <p>Your cart is empty. Please add items to your cart before submitting an order.</p>
        ) : (
          <>
            <Table columns={columns} dataSource={cartItems} pagination={{ pageSize: 5 }} />
            <div style={{ marginTop: 16, textAlign: 'right' }}>
              <strong>Total Price:</strong> ${totalPrice}
            </div>
          </>
        )}
      </Card>

      <Card title="Delivery Address" bordered={false} style={{ marginTop: 16 }}>
        <Form form={form} name="submit-order" onFinish={onFinish} layout="vertical" className="form-container">
          <Form.Item
            name="address_id"
            label="Select Address"
            rules={[{ required: true, message: 'Please select a delivery address!' }]}
          >
            <Select placeholder="Select an address" onChange={(value) => setSelectedAddress(value)}>
              {addresses.map((address) => (
                <Option key={address.id} value={address.id}>
                  {`${address.address_line1}, ${address.city}, ${address.state}, ${address.postal_code}, ${address.country}`}
                </Option>
              ))}
            </Select>
          </Form.Item>

          <Form.Item
            name="payment_method"
            label="Payment Method"
            rules={[{ required: true, message: 'Please select a payment method!' }]}
          >
            <Select placeholder="Select a payment method">
              <Option value="credit_card">Credit Card</Option>
              <Option value="debit_card">Debit Card</Option>
              <Option value="paypal">PayPal</Option>
              <Option value="bank_transfer">Bank Transfer</Option>
            </Select>
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              Submit Order
            </Button>
          </Form.Item>
          <Form.Item>
            <Link to="/cart/list">Back to Cart</Link>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}

export default OrderSubmitPage;