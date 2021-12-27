#include <windows.h>
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

int X = 0, Y = 0;
HDC hdc;
BOOL fVisible = FALSE;
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
        // centreaza totul
        OffsetRect(&rcCurrent, rc.right / 2, rc.bottom / 2);

        hdc = GetDC(hwnd);

        // seteaza orientarea
        SetViewportOrgEx(hdc, rcCurrent.left, rcCurrent.top, NULL);
        SetROP2(hdc, R2_NOT);

        return 0L;

    case WM_DESTROY:
        PostQuitMessage(0);
        return 0L;

    case WM_KEYDOWN:

        // ascunde patratul
        if (fVisible)
        {
            Rectangle(hdc, 100, 100, 10, 10);
        }

        int rect_size = 10;

        GetClientRect(hwnd, &rc);

        switch (wParam)
        {
        case VK_RIGHT:
            if (rcCurrent.right + X > rc.right)
            {
                MessageBox(hwnd, "Hey, you have reached the limit", "Warning", MB_OK);
                return 0L;
            }
            X = X + rect_size;
            break;

        case VK_LEFT:
            if (rcCurrent.left + X < rc.left)
            {
                MessageBox(hwnd, "Hey, you have reached the limit", "Warning", MB_OK);
                return 0L;
            }
            X = X - rect_size;
            break;

        case VK_UP:
            if (rcCurrent.top + Y < rc.top)
            {
                MessageBox(hwnd, "Hey, you have reached the limit", "Warning", MB_OK);
                return 0L;
            }
            Y = Y - rect_size;
            break;

        case VK_DOWN:
            if (rcCurrent.bottom + Y > rc.bottom)
            {
                MessageBox(hwnd, "Hey, you have reached the limit", "Warning", MB_OK);
                return 0L;
            }
            Y = Y + rect_size;
            break;

        default:
            break;
        }

        OffsetRect(&rcCurrent, X, Y);
        SetViewportOrgEx(hdc, rcCurrent.left, rcCurrent.top, NULL);
        fVisible = Rectangle(hdc, 100, 100, 10, 10);

        return 0L;

    case WM_PAINT:
        BeginPaint(hwnd, &ps);
        Ellipse(hdc, 200, 200, 50, 50);
        RoundRect(hdc, -250, 250, 50, 120, 50, 50); // 0 for angle

        if (!fVisible)
        {
            fVisible = Rectangle(hdc, 100, 100, 10, 10);
        }

        EndPaint(hwnd, &ps);
        return 0L;
    }
    return DefWindowProc(hwnd, message, wParam, lParam);
}
