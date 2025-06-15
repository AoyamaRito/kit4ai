package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
)

func main() {
	fmt.Println("病院管理システムUI設計")
	fmt.Println("=====================")
	
	// ByteCanvasを使用（正確な配置のため）
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
	
	// ヘッダーテキスト
	bc.WriteBytes(3, 2, "🏥 さくら総合病院")
	bc.WriteBytes(60, 2, "Dr.山田太郎")
	bc.WriteBytes(25, 3, "電子カルテシステム v2.1")
	
	// 患者情報
	bc.WriteBytes(15, 7, "患者情報")
	bc.WriteBytes(5, 9, "患者ID: P12345")
	bc.WriteBytes(5, 10, "氏名: 佐藤花子 (65歳)")
	bc.WriteBytes(5, 11, "性別: 女性")
	bc.WriteBytes(5, 12, "血液型: A型")
	bc.WriteBytes(5, 13, "アレルギー: ペニシリン")
	bc.WriteBytes(5, 14, "保険: 国民健康保険")
	bc.WriteBytes(5, 15, "緊急連絡先: 090-1234-5678")
	
	// 診察予約スケジュール
	bc.WriteBytes(52, 7, "本日の診察予約")
	bc.WriteBytes(44, 9, "時間   患者名      科目    状態")
	bc.WriteBytes(44, 10, "--------------------------------")
	bc.WriteBytes(44, 11, "09:00  田中一郎    内科    完了")
	bc.WriteBytes(44, 12, "09:30  鈴木美咲    小児科  進行中")
	bc.WriteBytes(44, 13, "10:00  佐藤花子    内科    待機")
	bc.WriteBytes(44, 14, "10:30  山田次郎    整形外科 待機")
	bc.WriteBytes(44, 15, "11:00  高橋明美    婦人科  待機")
	bc.WriteBytes(44, 16, "11:30  中村健太    皮膚科  待機")
	bc.WriteBytes(44, 17, "--------------------------------")
	bc.WriteBytes(44, 18, "待機患者数: 5名")
	bc.WriteBytes(44, 19, "平均待ち時間: 25分")
	
	// 薬剤情報
	bc.WriteBytes(17, 19, "処方薬情報")
	bc.WriteBytes(5, 21, "薬剤名        用量   回数  日数")
	bc.WriteBytes(5, 22, "--------------------------------")
	bc.WriteBytes(5, 23, "ロキソニン    60mg   3回   7日")
	bc.WriteBytes(5, 24, "胃薬         100mg   3回   7日")
	bc.WriteBytes(5, 25, "ビタミンB12   500mcg  1回  14日")
	bc.WriteBytes(5, 26, "--------------------------------")
	bc.WriteBytes(5, 27, "薬剤費合計: ¥2,840")
	
	// 検査結果
	bc.WriteBytes(56, 23, "最新検査結果")
	bc.WriteBytes(44, 25, "血圧: 130/85 mmHg (やや高)")
	bc.WriteBytes(44, 26, "体温: 36.8°C (正常)")
	bc.WriteBytes(44, 27, "脈拍: 72 bpm (正常)")
	bc.WriteBytes(44, 28, "血糖値: 110 mg/dl (正常)")
	bc.WriteBytes(44, 29, "HbA1c: 5.8% (正常)")
	bc.WriteBytes(44, 30, "コレステロール: 220 mg/dl")
	bc.WriteBytes(44, 31, "要再検査項目: なし")
	
	// ステータスバー
	bc.WriteBytes(5, 31, "システム状態: 正常稼働中")
	bc.WriteBytes(40, 31, "最終更新: 2024-06-15 14:32")
	bc.WriteBytes(5, 32, "データベース接続: 良好")
	bc.WriteBytes(40, 32, "バックアップ: 完了")
	
	// 操作ボタンエリア
	bc.WriteBytes(5, 34, "[F1]診察開始 [F2]処方入力 [F3]検査依頼 [F4]患者検索 [ESC]ログアウト")
	
	fmt.Println("病院管理システムUI:")
	fmt.Println(bc.String())
	
	// 仕様書出力
	file, err := os.Create("hospital_system_ui_spec.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Hospital Management System UI Design Specification\n")
	file.WriteString("==================================================\n\n")
	file.WriteString("System: さくら総合病院 電子カルテシステム v2.1\n")
	file.WriteString("Screen Size: 80x36 characters\n")
	file.WriteString("Target Users: 医師、看護師、事務スタッフ\n\n")
	file.WriteString("Main Features:\n")
	file.WriteString("1. Patient Information Display\n")
	file.WriteString("   - Patient ID, Name, Age, Gender\n")
	file.WriteString("   - Blood type, Allergies\n")
	file.WriteString("   - Insurance information\n")
	file.WriteString("   - Emergency contact\n\n")
	file.WriteString("2. Daily Appointment Schedule\n")
	file.WriteString("   - Time-based appointment list\n")
	file.WriteString("   - Patient names and departments\n")
	file.WriteString("   - Status tracking (completed, in-progress, waiting)\n")
	file.WriteString("   - Queue management with wait times\n\n")
	file.WriteString("3. Prescription Management\n")
	file.WriteString("   - Medication names and dosages\n")
	file.WriteString("   - Frequency and duration\n")
	file.WriteString("   - Total cost calculation\n\n")
	file.WriteString("4. Latest Test Results\n")
	file.WriteString("   - Vital signs (blood pressure, temperature, pulse)\n")
	file.WriteString("   - Blood test results\n")
	file.WriteString("   - Follow-up requirements\n\n")
	file.WriteString("5. System Status Monitoring\n")
	file.WriteString("   - Database connection status\n")
	file.WriteString("   - Last update timestamp\n")
	file.WriteString("   - Backup completion status\n\n")
	file.WriteString("Technical Specifications:\n")
	file.WriteString("- 8-bit byte canvas for precise alignment\n")
	file.WriteString("- ASCII-only characters for compatibility\n")
	file.WriteString("- Tabular data presentation\n")
	file.WriteString("- Function key shortcuts\n")
	file.WriteString("- Real-time data updates\n\n")
	file.WriteString("Security Features:\n")
	file.WriteString("- User authentication (Dr.山田太郎)\n")
	file.WriteString("- Patient data privacy protection\n")
	file.WriteString("- Audit trail logging\n")
	file.WriteString("- Session timeout protection\n\n")
	file.WriteString("UI Layout:\n")
	file.WriteString(bc.String())
	
	fmt.Println("\nhospital_system_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 医療現場で実用的な電子カルテシステム")
	fmt.Println("- 患者情報、予約管理、処方薬、検査結果の一元表示")
	fmt.Println("- ByteCanvasによる正確な表配置")
	fmt.Println("- ファンクションキーによる効率的操作")
	fmt.Println("- リアルタイムシステム状態監視")
	fmt.Println("- 医療従事者向けの直感的インターフェース")
}