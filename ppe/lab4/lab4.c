#include <windows.h>
#include <windowsx.h>
#include <wchar.h>
#include <stdlib.h>
#include <math.h>
#include <wingdi.h>

LRESULT WndProc(HWND hwnd, UINT message, WPARAM wParam, LPARAM lParam);

//Create window
int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int iCmdShow)
{

    static char szAppName[] = "Lab 2";
    HWND hwnd;
    MSG msg;
    WNDCLASSEX wndclass;
    wndclass.cbSize = sizeof(wndclass);
    wndclass.style = CS_HREDRAW | CS_VREDRAW;
    wndclass.lpfnWndProc = WndProc;
    wndclass.cbClsExtra = 0;
    wndclass.cbWndExtra = 0;
    wndclass.hInstance = hInstance;
    wndclass.hIcon = LoadIcon(NULL, IDI_APPLICATION);
    wndclass.hCursor = LoadCursor(NULL, IDC_ARROW);
    wndclass.hbrBackground = (HBRUSH)CreateSolidBrush(RGB(0, 128, 0)); // verde
    wndclass.lpszMenuName = NULL;
    wndclass.lpszClassName = szAppName;
    wndclass.hIconSm = LoadIcon(NULL, IDI_APPLICATION);
    RegisterClassEx(&wndclass);
    hwnd = CreateWindow(szAppName,                    // window class name
                        "Lucrarea de laborator nr.2", // window caption
                        WS_OVERLAPPEDWINDOW,          // window style
                        CW_USEDEFAULT,                // initial x position
                        CW_USEDEFAULT,                // initial y position
                        CW_USEDEFAULT,                // initial x size
                        CW_USEDEFAULT,                // initial y size
                        NULL,                         // parent window handle
                        NULL,                         // window menu handle
                        hInstance,                    // program instance handle
                        NULL);                        // creation parameters
    ShowWindow(hwnd, iCmdShow);
    UpdateWindow(hwnd);
    while (GetMessage(&msg, NULL, 0, 0))
    {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }
    return msg.wParam;
}

int X = 0, Y = 0, idTimer = -1;
int XEllipse = 0, YEllipse = 0;
int i = 0;
HDC hdc;
BOOL fVisible = FALSE;
BOOL fVisibleEllipse = FALSE;
RECT rcCurrent;

LRESULT WndProc(HWND hwnd, UINT message, WPARAM wParam, LPARAM lParam)
{
    static HBITMAP hBitmap;
    PAINTSTRUCT ps;
    RECT rc;

    switch (message)

    {
    case WM_CREATE:
        GetClientRect(hwnd, &rc);

        hdc = GetDC(hwnd);

        // seteaza orientarea
        SetViewportOrgEx(hdc, rcCurrent.left, rcCurrent.top, NULL);
        SetROP2(hdc, R2_NOT);

        return 0L;

    case WM_DESTROY:
        PostQuitMessage(0);
        KillTimer(hwnd, 1);
        return 0L;


    case WM_LBUTTONDOWN:
    {
        // incepe timer-ul care o sa genereze un loop.
        int xParam = GET_X_LPARAM(lParam);
        int yParam = GET_Y_LPARAM(lParam);
        Rectangle(hdc, xParam, yParam, xParam + 10, yParam + 10);

        SetColorAdjustment(hdc, WHITE_BRUSH);

        return 0L;
    }

    case WM_PAINT:
        BeginPaint(hwnd, &ps);

        EndPaint(hwnd, &ps);
        return 0L;
    }
    return DefWindowProc(hwnd, message, wParam, lParam);
}
