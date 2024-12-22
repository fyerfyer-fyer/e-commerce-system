import React from 'react';
import { Card, Form, Input, Button, Typography } from 'antd';

const { Title } = Typography;

function UserInfoPage() {
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('User info form values:', values);
    // Here you can add the logic to update user information on the backend
  };

  return (
    <div className="page-container">
      <Title level={2}>User Information</Title>
      <Card title="Edit User Information" bordered={false}>
        <Form form={form} name="user-info" onFinish={onFinish} layout="vertical" className="form-container">
          <Form.Item
            name="username"
            label="Username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input placeholder="Enter your username" />
          </Form.Item>
          <Form.Item
            name="email"
            label="Email"
            rules={[
              { required: true, message: 'Please enter your email!' },
              { type: 'email', message: 'Please enter a valid email address!' },
            ]}
          >
            <Input placeholder="Enter your email" />
          </Form.Item>
          <Form.Item
            name="phone"
            label="Phone Number"
            rules={[{ required: true, message: 'Please enter your phone number!' }]}
          >
            <Input placeholder="Enter your phone number" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              Save Changes
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
}

export default UserInfoPage;