import React from 'react';
import { Form, Input, Button, Typography, Rate } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

function CommentAddPage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Add comment form values:', values);
    // Here you can add the logic to add the comment on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Add Comment</Title>
      <Form form={form} name="add-comment" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="product_id"
          label="Product ID"
          rules={[{ required: true, message: 'Please enter the product ID!' }]}
        >
          <Input placeholder="Enter product ID" />
        </Form.Item>
        <Form.Item
          name="rating"
          label="Rating"
          rules={[{ required: true, message: 'Please select a rating!' }]}
        >
          <Rate allowHalf />
        </Form.Item>
        <Form.Item
          name="content"
          label="Comment"
          rules={[{ required: true, message: 'Please enter your comment!' }]}
        >
          <Input.TextArea rows={4} placeholder="Enter your comment" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Add Comment
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/comment/list/1">View Comments</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default CommentAddPage;