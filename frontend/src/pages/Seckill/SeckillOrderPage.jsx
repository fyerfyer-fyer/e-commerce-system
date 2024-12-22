import React, { useState } from 'react';
import { Form, Input, Button, Typography, Select, Space, InputNumber, Card } from 'antd';
import { useParams, Link } from 'react-router-dom';

const { Title } = Typography;
const { Option } = Select;

function SeckillOrderPage() {
  const { productId, eventId } = useParams();
  const [form] = Form.useForm();
  const [quantity, setQuantity] = useState(1);

  const onFinish = (values) => {
    console.log('Place seckill order form values:', values);
    // Here you can add the logic to place the seckill order on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Place Seckill Order</Title>
      <Card title={`Seckill Order for Product ${productId}`} bordered={false}>
        <Form form={form} name="place-seckill-order" onFinish={onFinish} layout="vertical" className="form-container">
          <Form.Item
            name="product_id"
            label="Product ID"
            initialValue={productId}
            hidden
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="event_id"
            label="Event ID"
            initialValue={eventId}
            hidden
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[{ required: true, message: 'Please enter the quantity!' }]}
          >
            <InputNumber min={1} max={10} defaultValue={1} onChange={(value) => setQuantity(value)} />
          </Form.Item>
          <Form.Item
            name="address_id"
            label="Select Address"
            rules={[{ required: true, message: 'Please select an address!' }]}
          >
            <Select placeholder="Select an address">
              <Option value="1">123 Main St, New York, NY</Option>
              <Option value="2">456 Elm St, Los Angeles, CA</Option>
            </Select>
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              Place Order
            </Button>
          </Form.Item>
          <Form.Item>
            <Link to={`/seckill/products?eventId=${eventId}`}>Back to Products</Link>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}

export default SeckillOrderPage;