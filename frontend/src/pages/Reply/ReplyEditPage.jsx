import React, { useEffect } from 'react';
import { Form, Input, Button, Typography } from 'antd';
import { useParams } from 'react-router-dom';

const { Title } = Typography;

function ReplyEditPage() {
  const { replyId } = useParams();
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Edit reply form values:', values);
    // Here you can add the logic to edit the reply on the backend
  };

  // Simulate fetching reply data from the backend
  useEffect(() => {
    // Simulate an API call to fetch the reply data
    setTimeout(() => {
      form.setFieldsValue({
        content: 'This is a great comment!',
      });
    }, 1000);
  }, [replyId, form]);

  return (
    <div className="page-container">
      <Title level={2}>Edit Reply</Title>
      <Form form={form} name="edit-reply" onFinish={onFinish} layout="vertical" className="form-container">
        <Form.Item
          name="content"
          label="Reply"
          rules={[{ required: true, message: 'Please enter your reply!' }]}
        >
          <Input.TextArea rows={4} placeholder="Enter your reply" />
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
      </Form>
    </div>
  );
}

export default ReplyEditPage;