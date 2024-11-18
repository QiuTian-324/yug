import { Message } from '@arco-design/web-vue';


export const newInfoMessage = (text: string, duration: number = 2000, show = true) => {
  if (show) {
    Message["info"]({
      content: text,
      duration,
      closable: true
    });
  }
};

export const newSuccessMessage = (text: string, duration: number = 2000, show = true) => {
  if (show) {
    Message["success"]({
      content: text,
      duration,
      closable: true
    });
  }
};
export const newErrorMessage = (text: string, duration: number = 2000, show = true) => {
  if (show) {
    Message["error"]({
      content: text,
      duration,
      closable: true
    });
  }
};

export const newWarningMessage = (text: string, duration: number = 2000, show = true) => {
  if (show) {
    Message["warning"]({
      content: text,
      duration,
      closable: true
    });
  }
};

export const newLoadingMessage = (text: string, duration: number, show = true) => {
  if (show) {
    Message["loading"]({
      content: text,
      duration,
      closable: true
    });
  }
};

export const newNormalMessage = (text: string, duration: number, show = true) => {
  if (show) {
    Message["info"]({
      content: text,
      duration,
      closable: true
    });
  }
};



// const showMessage = Symbol('showMessage');

// // 封装message类
// class MessageBox {
//   success(text: string, duration: number, show = true): void {
//     this[showMessage]('success', text, duration, show);
//   }

//   error(text: string, duration: number, show = true): void {
//     this[showMessage]('error', text, duration, show);
//   }

//   warning(text: string, duration: number, show = true): void {
//     this[showMessage]('warning', text, duration, show);
//   }

//   info(text: string, duration: number, show = true): void {
//     this[showMessage]('info', text, duration, show);
//   }

//   loading(text: string, duration: number, show = true): void {
//     this[showMessage]('loading', text, duration, show);
//   }

//   normal(text: string, duration: number, show = true): void {
//     this[showMessage]('info', text, duration, show); // 使用 info 作为 normal 的实现
//   }

//   clear(): void {
//     // 清除所有消息
//     const messages = document.getElementsByClassName('arco-message');
//     for (let i = 0; i < messages.length; i++) {
//       messages[i].remove();
//     }
//   }

//   [showMessage](type: string, text: string, duration: number, show: boolean) {
//     if (show) {
//       Message[type]({
//         content: text,
//         duration,
//         closable: true
//       });
//     } else {
//       Message[type]({
//         content: text,
//         duration,
//         closable: true
//       });
//     }
//   }
// }

// export default new MessageBox();