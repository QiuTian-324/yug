import { electronApp, is, optimizer } from '@electron-toolkit/utils';
import { app, BrowserWindow, ipcMain, shell } from 'electron';
import { join } from 'path';
import icon from '../../resources/icon.png?asset';
// 声明 mainWindow 为 BrowserWindow | null 类型
let mainWindow: BrowserWindow | null = null;
function createWindow(): void {
  // 初始化 mainWindow
  mainWindow = new BrowserWindow({
    width: 950,
    height: 670,
    show: false,
    autoHideMenuBar: true,
    fullscreenable: true, // 禁用全屏按钮
    maximizable: false,    // 启用最大化
    resizable: true,      // 启用调整大小
    ...(process.platform === 'linux' ? { icon } : {}),
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false
    }
  });

  mainWindow.on('ready-to-show', () => {
    mainWindow?.show();
  });

  mainWindow.on('closed', () => {
    mainWindow = null;
  });

  mainWindow.webContents.setWindowOpenHandler((details) => {
    shell.openExternal(details.url);
    return { action: 'deny' };
  });

  if (is.dev && process.env['ELECTRON_RENDERER_URL']) {
    mainWindow.loadURL(process.env['ELECTRON_RENDERER_URL']);
  } else {
    mainWindow.loadFile(join(__dirname, '../renderer/index.html'));
  }
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
  // Set app user model id for windows
  electronApp.setAppUserModelId('com.electron')

  // Default open or close DevTools by F12 in development
  // and ignore CommandOrControl + R in production.
  // see https://github.com/alex8088/electron-toolkit/tree/master/packages/utils
  app.on('browser-window-created', (_, window) => {
    optimizer.watchWindowShortcuts(window)
  })

  ipcMain.on('isLogin', (event, token) => {
    console.log('token:', token);
    if (mainWindow) {
      if (!token) {
        mainWindow.setSize(320, 450);
        mainWindow.setMaximizable(false); // 禁用最大化
        mainWindow.setResizable(false);   // 禁用调整大小
      } else {
        mainWindow.setSize(950, 670);
        mainWindow.setMaximizable(true);  // 启用最大化
        mainWindow.setResizable(true);    // 启用调整大小
      }
    }

    // 发送 'pong' 消息回到渲染进程
    event.sender.send('pong', 'Hello from main process');
  });

  createWindow()

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

// In this file you can include the rest of your app"s specific main process
// code. You can also put them in separate files and require them here.
