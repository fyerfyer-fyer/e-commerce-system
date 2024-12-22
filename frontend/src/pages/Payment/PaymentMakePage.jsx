import React from 'react';
import { Form, Input, Button, Typography, Select, Space, InputNumber} from 'antd';
import { Link} from 'react-router-dom';

const { Title } = Typography;
const { Option } = Select;

function PaymentMakePage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Make payment form values:', values);
    // Here you can add the logic to make the payment on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Make Payment</Title>
      <Form form={form} name="make-payment" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="order_id"
          label="Order ID"
          rules={[{ required: true, message: 'Please enter the order ID!' }]}
        >
          <Input placeholder="Enter order ID" />
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
            <Option value="wallet">Wallet</Option>
          </Select>
        </Form.Item>
        <Form.Item
          name="amount"
          label="Amount"
          rules={[{ required: true, message: 'Please enter the payment amount!' }]}
        >
          <InputNumber min={0.01} precision={2} placeholder="Enter payment amount" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Make Payment
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/order/list">Back to Orders</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default PaymentMakePage;