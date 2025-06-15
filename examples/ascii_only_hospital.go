package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
)

func main() {
	fmt.Println("ASCII Only Hospital System UI")
	fmt.Println("=============================")
	
	// ByteCanvasを使用（ASCII文字のみ）
	bc := canvas.NewByteCanvas()
	
	// メインフレーム
	bc.DrawBox(0, 0, 79, 35)
	
	// ヘッダーエリア
	bc.DrawBox(1, 1, 78, 4)
	
	// 患者情報カード
	bc.DrawBox(3, 6, 38, 16)
	
	// 診察予約スケジュール
	bc.DrawBox(42, 6, 77, 20)
	
	// 薬剤情報
	bc.DrawBox(3, 18, 38, 28)
	
	// 検査結果
	bc.DrawBox(42, 22, 77, 32)
	
	// ステータスバー
	bc.DrawBox(3, 30, 77, 33)
	
	// ヘッダーテキスト（ASCII only）
	bc.WriteBytes(3, 2, "SAKURA GENERAL HOSPITAL")
	bc.WriteBytes(60, 2, "Dr.YAMADA")
	bc.WriteBytes(25, 3, "Electronic Medical Records v2.1")
	
	// 患者情報（英語表記）
	bc.WriteBytes(15, 7, "PATIENT INFO")
	bc.WriteBytes(5, 9, "Patient ID: P12345")
	bc.WriteBytes(5, 10, "Name: SATO HANAKO (65F)")
	bc.WriteBytes(5, 11, "Gender: Female")
	bc.WriteBytes(5, 12, "Blood Type: A")
	bc.WriteBytes(5, 13, "Allergy: Penicillin")
	bc.WriteBytes(5, 14, "Insurance: National Health")
	bc.WriteBytes(5, 15, "Emergency: 090-1234-5678")
	
	// 診察予約スケジュール
	bc.WriteBytes(52, 7, "TODAY'S SCHEDULE")
	bc.WriteBytes(44, 9, "Time   Patient      Dept     Status")
	bc.WriteBytes(44, 10, "----------------------------------")
	bc.WriteBytes(44, 11, "09:00  TANAKA T.    Internal Done")
	bc.WriteBytes(44, 12, "09:30  SUZUKI M.    Pediatr  Active")
	bc.WriteBytes(44, 13, "10:00  SATO H.      Internal Wait")
	bc.WriteBytes(44, 14, "10:30  YAMADA J.    Orthoped Wait")
	bc.WriteBytes(44, 15, "11:00  TAKAHASHI A. Gynecol  Wait")
	bc.WriteBytes(44, 16, "11:30  NAKAMURA K.  Dermatol Wait")
	bc.WriteBytes(44, 17, "----------------------------------")
	bc.WriteBytes(44, 18, "Waiting Patients: 5")
	bc.WriteBytes(44, 19, "Average Wait: 25min")
	
	// 薬剤情報
	bc.WriteBytes(15, 19, "PRESCRIPTION")
	bc.WriteBytes(5, 21, "Drug Name      Dose    Freq  Days")
	bc.WriteBytes(5, 22, "----------------------------------")
	bc.WriteBytes(5, 23, "Loxonin        60mg    3x    7")
	bc.WriteBytes(5, 24, "Gastric Med    100mg   3x    7")
	bc.WriteBytes(5, 25, "Vitamin B12    500mcg  1x    14")
	bc.WriteBytes(5, 26, "----------------------------------")
	bc.WriteBytes(5, 27, "Total Cost: $28.40")
	
	// 検査結果
	bc.WriteBytes(54, 23, "LATEST TEST RESULTS")
	bc.WriteBytes(44, 25, "Blood Pressure: 130/85 mmHg (High)")
	bc.WriteBytes(44, 26, "Temperature: 36.8C (Normal)")
	bc.WriteBytes(44, 27, "Pulse: 72 bpm (Normal)")
	bc.WriteBytes(44, 28, "Blood Sugar: 110 mg/dl (Normal)")
	bc.WriteBytes(44, 29, "HbA1c: 5.8% (Normal)")
	bc.WriteBytes(44, 30, "Cholesterol: 220 mg/dl")
	bc.WriteBytes(44, 31, "Follow-up: None required")
	
	// ステータスバー
	bc.WriteBytes(5, 31, "System Status: Normal Operation")
	bc.WriteBytes(40, 31, "Last Update: 2024-06-15 14:32")
	bc.WriteBytes(5, 32, "Database: Connected")
	bc.WriteBytes(40, 32, "Backup: Completed")
	
	// 操作ボタンエリア
	bc.WriteBytes(5, 34, "[F1]Start [F2]Prescribe [F3]Lab Test [F4]Search [ESC]Logout")
	
	fmt.Println("ASCII Only Hospital System UI:")
	fmt.Println(bc.String())
	
	// ASCII安全な仕様書出力
	file, err := os.Create("ascii_hospital_ui_spec.txt")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("ASCII-Safe Hospital Management System UI\n")
	file.WriteString("========================================\n\n")
	file.WriteString("System: SAKURA GENERAL HOSPITAL EMR v2.1\n")
	file.WriteString("Screen Size: 80x36 characters\n")
	file.WriteString("Character Set: ASCII only (compatibility guaranteed)\n")
	file.WriteString("Target Users: Medical staff (doctors, nurses, admin)\n\n")
	file.WriteString("Design Principles:\n")
	file.WriteString("- ASCII-only characters for universal compatibility\n")
	file.WriteString("- 8-bit byte canvas for precise alignment\n")
	file.WriteString("- Tabular data layout for medical information\n")
	file.WriteString("- Function key shortcuts for efficiency\n")
	file.WriteString("- Real-time status monitoring\n\n")
	file.WriteString("Main Components:\n")
	file.WriteString("1. Header - Hospital name, doctor, system version\n")
	file.WriteString("2. Patient Information - ID, demographics, medical info\n")
	file.WriteString("3. Daily Schedule - Appointment queue with status\n")
	file.WriteString("4. Prescription Management - Drug list with dosages\n")
	file.WriteString("5. Test Results - Latest lab values and vitals\n")
	file.WriteString("6. System Status - Database and backup monitoring\n")
	file.WriteString("7. Function Keys - Quick access to main operations\n\n")
	file.WriteString("Compatibility Features:\n")
	file.WriteString("- Works in any text editor or terminal\n")
	file.WriteString("- No special fonts or Unicode required\n")
	file.WriteString("- Printable on any ASCII-compatible printer\n")
	file.WriteString("- Network-safe transmission\n")
	file.WriteString("- Legacy system compatible\n\n")
	file.WriteString("UI Layout:\n")
	file.WriteString(bc.String())
	
	fmt.Println("\nascii_hospital_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 完全ASCII文字のみ使用")
	fmt.Println("- どのテキストエディタでも正確に表示")
	fmt.Println("- レガシーシステム対応")
	fmt.Println("- ネットワーク安全な文字セット")
	fmt.Println("- プリンタ出力対応")
}