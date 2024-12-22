import React from 'react';
import { Form, Input, Button, Typography, Select, InputNumber } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;
const { Option } = Select;

function CartAddPage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Add to cart form values:', values);
    // Here you can add the logic to add the product to the cart on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Add to Cart</Title>
      <Form form={form} name="add-to-cart" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="product_id"
          label="Product ID"
          rules={[{ required: true, message: 'Please enter the product ID!' }]}
        >
          <Input placeholder="Enter product ID" />
        </Form.Item>
        <Form.Item
          name="quantity"
          label="Quantity"
          rules={[{ required: true, message: 'Please enter the quantity!' }]}
        >
          <InputNumber min={1} defaultValue={1} />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Add to Cart
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/cart/list">View Cart</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default CartAddPage;