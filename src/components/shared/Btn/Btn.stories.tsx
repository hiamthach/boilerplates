import type { Meta, StoryObj } from '@storybook/react';

import Btn from './Btn';

const meta = {
  title: 'Shared/Btn',
  component: Btn,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Btn>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: 'Default',
    type: 'primary',
  },
};

export const Outline: Story = {
  args: {
    children: 'Outline',
    type: 'default',
  },
};

export const Large: Story = {
  args: {
    children: 'Large',
    size: 'large',
  },
};

export const Small: Story = {
  args: {
    children: 'Small',
    size: 'small',
  },
};
