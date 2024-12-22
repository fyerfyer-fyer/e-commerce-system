import React, { useEffect } from 'react';
import { Form, Input, Button, Typography, Select, Space, Upload, message, InputNumber } from 'antd';
import { Link, useParams } from 'react-router-dom';
import { InboxOutlined } from '@ant-design/icons';

const { Title } = Typography;
const { Option } = Select;

function ProductEditPage() {
  const { productId } = useParams();
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Edit product form values:', values);
    // Here you can add the logic to edit the product on the backend
    message.success('Product updated successfully!');
  };

  // Simulate fetching product data from the backend
  useEffect(() => {
    // Simulate an API call to fetch the product data
    setTimeout(() => {
      form.setFieldsValue({
        name: 'Smartphone',
        description: 'A high-quality smartphone with advanced features.',
        price: 499.99,
        stock: 100,
        category: 'electronics',
        image: 'https://example.com/image.jpg',
      });
    }, 1000);
  }, [productId, form]);

  const normFile = (e) => {
    if (Array.isArray(e)) {
      return e;
    }
    return e && e.fileList;
  };

  return (
    <div className="page-container">
      <Title level={2}>Edit Product</Title>
      <Form form={form} name="edit-product" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="name"
          label="Product Name"
          rules={[{ required: true, message: 'Please enter the product name!' }]}
        >
          <Input placeholder="Enter product name" />
        </Form.Item>
        <Form.Item
          name="description"
          label="Description"
          rules={[{ required: true, message: 'Please enter the product description!' }]}
        >
          <Input.TextArea rows={4} placeholder="Enter product description" />
        </Form.Item>
        <Form.Item
          name="price"
          label="Price"
          rules={[{ required: true, message: 'Please enter the product price!' }]}
        >
          <InputNumber min={0.01} precision={2} placeholder="Enter product price" />
        </Form.Item>
        <Form.Item
          name="stock"
          label="Stock"
          rules={[{ required: true, message: 'Please enter the product stock!' }]}
        >
          <InputNumber min={0} placeholder="Enter product stock" />
        </Form.Item>
        <Form.Item
          name="category"
          label="Category"
          rules={[{ required: true, message: 'Please select a category!' }]}
        >
          <Select placeholder="Select a category">
            <Option value="electronics">Electronics</Option>
            <Option value="fashion">Fashion</Option>
            <Option value="home">Home & Garden</Option>
            <Option value="sports">Sports & Outdoors</Option>
          </Select>
        </Form.Item>
        <Form.Item
          name="image"
          label="Product Image"
          valuePropName="fileList"
          getValueFromEvent={normFile}
          rules={[{ required: true, message: 'Please upload a product image!' }]}
        >
          <Upload.Dragger name="files" action="/upload" listType="picture">
            <p className="ant-upload-drag-icon">
              <InboxOutlined />
            </p>
            <p className="ant-upload-text">Click or drag file to this area to upload</p>
            <p className="ant-upload-hint">Support for a single or bulk upload.</p>
          </Upload.Dragger>
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Save Changes
          </Button>
        </Form.Item>
        <Form.Item>
          <Button type="link" onClick={() => form.resetFields()}>
            Reset
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/admin/product/list">Back to Products</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default ProductEditPage;
