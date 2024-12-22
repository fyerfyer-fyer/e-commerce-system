import React from 'react';
import { Form, Input, Button, Typography, Space, InputNumber } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

function PaymentRefundPage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Refund request form values:', values);
    // Here you can add the logic to request a refund on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Request Refund</Title>
      <Form form={form} name="refund-request" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="payment_id"
          label="Payment ID"
          rules={[{ required: true, message: 'Please enter the payment ID!' }]}
        >
          <Input placeholder="Enter payment ID" />
        </Form.Item>
        <Form.Item
          name="refund_amount"
          label="Refund Amount"
          rules={[{ required: true, message: 'Please enter the refund amount!' }]}
        >
          <InputNumber min={0.01} precision={2} placeholder="Enter refund amount" />
        </Form.Item>
        <Form.Item
          name="refund_reason"
          label="Reason for Refund"
          rules={[{ required: true, message: 'Please enter the reason for refund!' }]}
        >
          <Input.TextArea rows={4} placeholder="Enter reason for refund" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Request Refund
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/payment/status/1">Check Payment Status</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default PaymentRefundPage;