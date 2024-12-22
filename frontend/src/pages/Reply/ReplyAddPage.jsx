import React from 'react';
import { Form, Input, Button, Typography } from 'antd';
import { Link } from 'react-router-dom';

const { Title } = Typography;

function ReplyAddPage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Add reply form values:', values);
    // Here you can add the logic to add the reply on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>Add Reply</Title>
      <Form form={form} name="add-reply" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="comment_id"
          label="Comment ID"
          rules={[{ required: true, message: 'Please enter the comment ID!' }]}
        >
          <Input placeholder="Enter comment ID" />
        </Form.Item>
        <Form.Item
          name="content"
          label="Reply"
          rules={[{ required: true, message: 'Please enter your reply!' }]}
        >
          <Input.TextArea rows={4} placeholder="Enter your reply" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Add Reply
          </Button>
        </Form.Item>
        <Form.Item>
          <Link to="/reply/list/1">View Replies</Link>
        </Form.Item>
      </Form>
    </div>
  );
}

export default ReplyAddPage;