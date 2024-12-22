import React, { useEffect } from 'react';
import { Form, Input, Button, Typography, Rate } from 'antd';
import { useParams } from 'react-router-dom';

const { Title } = Typography;

function CommentEditPage() {
  const { commentId } = useParams();
  const [form] = Form.useForm();

  const onFinish = (values) => {
    console.log('Edit comment form values:', values);
    // Here you can add the logic to edit the comment on the backend
  };

  // Simulate fetching comment data from the backend
  useEffect(() => {
    // Simulate an API call to fetch the comment data
    setTimeout(() => {
      form.setFieldsValue({
        rating: 4,
        content: 'This is a great product!',
      });
    }, 1000);
  }, [commentId, form]);

  return (
    <div className="page-container">
      <Title level={2}>Edit Comment</Title>
      <Form form={form} name="edit-comment" onFinish={onFinish} layout="vertical" className="form-container">
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

export default CommentEditPage;