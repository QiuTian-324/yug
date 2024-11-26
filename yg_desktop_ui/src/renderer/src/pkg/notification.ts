import { notification } from 'ant-design-vue';

export const newInfoNotification = (message: string, description: string, duration: number = 2000, show = true) => {
  if (show) {
    notification["info"]({
      message,
      description,
      duration,
    });
  }
};

export const newSuccessNotification = (message: string, description: string, duration: number = 2000, show = true) => {
  if (show) {
    notification["success"]({
      message,
      description,
      duration,
    });
  }
};

export const newErrorNotification = (message: string, description: string, duration: number = 2000, show = true) => {
  if (show) {
    notification["error"]({
      message,
      description,
      duration,
    });
  }
};

export const newWarningNotification = (message: string, description: string, duration: number = 2000, show = true) => {
  if (show) {
    notification["warning"]({
      message,
      description,
      duration,
    });
  }
};