import { message } from 'ant-design-vue';

export const newInfoMessage = (text: string, duration: number = 3, show = true) => {
  if (show) {
    message["info"]({
      content: text,
      duration,
    });
  }
};

export const newSuccessMessage = (text: string, duration: number = 3, show = true) => {
  if (show) {
    message["success"]({
      content: text,
      duration,
    });
  }
};

export const newErrorMessage = (text: string, duration: number = 3, show = true) => {
  if (show) {
    message["error"]({
      content: text,
      duration,
    });
  }
};

export const newWarningMessage = (text: string, duration: number = 3, show = true) => {
  if (show) {
    message["warning"]({
      content: text,
      duration,
    });
  }
};

export const newLoadingMessage = (text: string, duration: number, show = true) => {
  if (show) {
    message["loading"]({
      content: text,
      duration,
    });
  }
};

export const newNormalMessage = (text: string, duration: number, show = true) => {
  if (show) {
    message["info"]({
      content: text,
      duration,
    });
  }
};
