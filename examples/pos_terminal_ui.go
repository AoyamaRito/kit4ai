package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
)

func main() {
	fmt.Println("POS Terminal System UI Design")
	fmt.Println("============================")
	
	// ByteCanvasを使用（ASCII安全）
	bc := canvas.NewByteCanvas()
	
	// メインフレーム
	bc.DrawBox(0, 0, 79, 40)
	
	// ヘッダーエリア
	bc.DrawBox(1, 1, 78, 4)
	
	// 商品一覧エリア
	bc.DrawBox(3, 6, 50, 25)
	
	// 小計・計算エリア
	bc.DrawBox(52, 6, 77, 18)
	
	// 支払い方法エリア
	bc.DrawBox(52, 20, 77, 32)
	
	// 操作ボタンエリア
	bc.DrawBox(3, 27, 50, 32)
	
	// ステータスバー
	bc.DrawBox(3, 34, 77, 38)
	
	// ヘッダー情報
	bc.WriteBytes(3, 2, "KONBINI STORE POS SYSTEM v3.2")
	bc.WriteBytes(60, 2, "Terminal: 001")
	bc.WriteBytes(25, 3, "Cashier: TANAKA YUKI")
	bc.WriteBytes(60, 3, "2024-06-15 14:45")
	
	// 商品一覧ヘッダー
	bc.WriteBytes(18, 7, "SALES TRANSACTION")
	bc.WriteBytes(5, 9, "Item Code    Description        Qty  Price  Total")
	bc.WriteBytes(5, 10, "------------------------------------------------")
	
	// 商品リスト
	bc.WriteBytes(5, 11, "4901234567890 Coca Cola 500ml      2   150    300")
	bc.WriteBytes(5, 12, "4987654321098 Onigiri Salmon       1   180    180")
	bc.WriteBytes(5, 13, "4912345678901 Potato Chips         1   120    120")
	bc.WriteBytes(5, 14, "4956789012345 Green Tea 350ml      1   110    110")
	bc.WriteBytes(5, 15, "4923456789012 Sandwich Ham&Egg     1   250    250")
	bc.WriteBytes(5, 16, "4934567890123 Cigarette Pack       1   580    580")
	bc.WriteBytes(5, 17, "4945678901234 Weekly Magazine      1   400    400")
	bc.WriteBytes(5, 18, "4978901234567 AA Batteries 4pk     1   320    320")
	bc.WriteBytes(5, 19, "------------------------------------------------")
	bc.WriteBytes(5, 20, "                              8 items  Subtotal:")
	bc.WriteBytes(5, 21, "")
	bc.WriteBytes(5, 22, "                                    Tax (10%):")
	bc.WriteBytes(5, 23, "                                       TOTAL:")
	
	// 小計・税込計算
	bc.WriteBytes(62, 7, "CALCULATION")
	bc.WriteBytes(54, 9, "Subtotal:    2,260")
	bc.WriteBytes(54, 10, "Tax (10%):     226")
	bc.WriteBytes(54, 11, "----------")
	bc.WriteBytes(54, 12, "TOTAL:     2,486")
	bc.WriteBytes(54, 14, "Points Used:     0")
	bc.WriteBytes(54, 15, "Discount:        0")
	bc.WriteBytes(54, 16, "----------")
	bc.WriteBytes(54, 17, "Final Total: 2,486")
	
	// 支払い方法
	bc.WriteBytes(60, 21, "PAYMENT METHOD")
	bc.WriteBytes(54, 23, "[ ] Cash")
	bc.WriteBytes(54, 24, "[X] Credit Card")
	bc.WriteBytes(54, 25, "[ ] IC Card")
	bc.WriteBytes(54, 26, "[ ] QR Code")
	bc.WriteBytes(54, 27, "[ ] Points")
	bc.WriteBytes(54, 29, "Card: **** 1234")
	bc.WriteBytes(54, 30, "Auth: OK")
	bc.WriteBytes(54, 31, "Receipt: Print")
	
	// 操作ボタン
	bc.WriteBytes(18, 28, "OPERATIONS")
	bc.WriteBytes(5, 30, "[F1]Add Item [F2]Delete [F3]Discount [F4]Points")
	bc.WriteBytes(5, 31, "[F5]Payment  [F6]Cancel  [F8]Manager [ESC]Logout")
	
	// ステータス情報
	bc.WriteBytes(5, 35, "Status: Transaction in Progress")
	bc.WriteBytes(5, 36, "Card Reader: Connected")
	bc.WriteBytes(40, 35, "Printer: Ready")
	bc.WriteBytes(40, 36, "Cash Drawer: Closed")
	bc.WriteBytes(5, 37, "Network: Online")
	bc.WriteBytes(40, 37, "Last Backup: 14:30")
	
	fmt.Println("POS Terminal UI:")
	fmt.Println(bc.String())
	
	// 仕様書出力
	file, err := os.Create("pos_terminal_ui_spec.txt")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("POS Terminal System UI Design Specification\n")
	file.WriteString("===========================================\n\n")
	file.WriteString("System: KONBINI STORE POS SYSTEM v3.2\n")
	file.WriteString("Terminal Type: Retail Point of Sale\n")
	file.WriteString("Screen Size: 80x41 characters\n")
	file.WriteString("Target Users: Store cashiers and managers\n\n")
	file.WriteString("Core Functions:\n")
	file.WriteString("1. Product Scanning and Entry\n")
	file.WriteString("   - Barcode scanning support\n")
	file.WriteString("   - Manual product code entry\n")
	file.WriteString("   - Quantity adjustment\n")
	file.WriteString("   - Price lookup and calculation\n\n")
	file.WriteString("2. Transaction Management\n")
	file.WriteString("   - Real-time subtotal calculation\n")
	file.WriteString("   - Automatic tax calculation (10%)\n")
	file.WriteString("   - Discount and promotion handling\n")
	file.WriteString("   - Customer loyalty points integration\n\n")
	file.WriteString("3. Payment Processing\n")
	file.WriteString("   - Multiple payment methods support\n")
	file.WriteString("     * Cash payments\n")
	file.WriteString("     * Credit/Debit cards\n")
	file.WriteString("     * IC cards (contactless)\n")
	file.WriteString("     * QR code payments\n")
	file.WriteString("     * Loyalty points redemption\n")
	file.WriteString("   - Card authorization processing\n")
	file.WriteString("   - Receipt printing control\n\n")
	file.WriteString("4. System Integration\n")
	file.WriteString("   - Inventory management sync\n")
	file.WriteString("   - Sales reporting\n")
	file.WriteString("   - Network connectivity monitoring\n")
	file.WriteString("   - Hardware status tracking\n\n")
	file.WriteString("Hardware Components:\n")
	file.WriteString("- Barcode scanner\n")
	file.WriteString("- Card reader (magnetic stripe/chip/contactless)\n")
	file.WriteString("- Receipt printer\n")
	file.WriteString("- Cash drawer\n")
	file.WriteString("- Customer display\n")
	file.WriteString("- Network connection\n\n")
	file.WriteString("Function Key Operations:\n")
	file.WriteString("F1 - Add Item (manual entry)\n")
	file.WriteString("F2 - Delete Item (remove from transaction)\n")
	file.WriteString("F3 - Apply Discount\n")
	file.WriteString("F4 - Points Management\n")
	file.WriteString("F5 - Payment Processing\n")
	file.WriteString("F6 - Cancel Transaction\n")
	file.WriteString("F8 - Manager Functions\n")
	file.WriteString("ESC - Logout/Exit\n\n")
	file.WriteString("Security Features:\n")
	file.WriteString("- Cashier authentication\n")
	file.WriteString("- Manager authorization for special functions\n")
	file.WriteString("- Transaction logging\n")
	file.WriteString("- Secure payment processing\n")
	file.WriteString("- Data backup scheduling\n\n")
	file.WriteString("UI Layout:\n")
	file.WriteString(bc.String())
	
	fmt.Println("\npos_terminal_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 実用的なコンビニPOSシステム")
	fmt.Println("- 商品スキャン、計算、決済の完全フロー")
	fmt.Println("- 複数決済方法対応（現金、カード、IC、QR）")
	fmt.Println("- リアルタイム税計算と小計表示")
	fmt.Println("- ハードウェア状態監視")
	fmt.Println("- ファンクションキーによる効率操作")
	fmt.Println("- ASCII安全で安定表示")
}