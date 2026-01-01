import { Assessment, getAssessments } from "@/api/assessment.api";
import { getClasses } from "@/api/class.api";
import { Exam, getStudentClassExams } from "@/api/exam.api";
import { useAuth } from "@/context/AuthContext";
import {
  AlertCircle,
  Award,
  CheckCircle2,
  Clock,
  FileText,
  Plus,
} from "lucide-react-native";
import { useEffect, useState } from "react";
import {
  ActivityIndicator,
  Alert,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";

interface AssessmentWithClass extends Assessment {
  className: string;
  classId: string;
}

interface ExamWithClass extends Exam {
  className: string;
  classId: string;
}

type QuizItem = (AssessmentWithClass | ExamWithClass) & {
  type: "assessment" | "exam";
};

export default function QuizScreen() {
  const { user } = useAuth();
  const isStudent = user?.role?.name?.toLowerCase() === "siswa";
  const [selectedTab, setSelectedTab] = useState("all");
  const [loading, setLoading] = useState(true);
  const [quizItems, setQuizItems] = useState<QuizItem[]>([]);

  const fetchQuizData = async () => {
    try {
      setLoading(true);
      const classesResponse = await getClasses();
      const allItems: QuizItem[] = [];

      for (const classItem of classesResponse.data) {
        try {
          const assessmentResponse = await getAssessments(classItem.id);
          const assessmentsWithClass = assessmentResponse.data.data.map(
            (assessment: Assessment) => ({
              ...assessment,
              className: classItem.name,
              classId: classItem.id,
              type: "assessment" as const,
            })
          );
          allItems.push(...assessmentsWithClass);

          if (isStudent) {
            try {
              const examResponse = await getStudentClassExams(classItem.id);
              const examsWithClass = examResponse.data.data.map(
                (exam: Exam) => ({
                  ...exam,
                  className: classItem.name,
                  classId: classItem.id,
                  type: "exam" as const,
                })
              );
              allItems.push(...examsWithClass);
            } catch {
              console.log(`Failed to fetch exams for class ${classItem.id}`);
            }
          }
        } catch {
          console.log(`Failed to fetch quiz data for class ${classItem.id}`);
        }
      }

      setQuizItems(allItems);
    } catch (error: any) {
      Alert.alert(
        "Error",
        error?.response?.data?.message || "Gagal ambil data"
      );
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchQuizData();
  }, []);

  const filterQuizItems = () => {
    if (selectedTab === "all") return quizItems;
    if (selectedTab === "exam")
      return quizItems.filter((item) => item.type === "exam");
    return quizItems.filter((item) => item.type === "assessment");
  };

  const filteredItems = filterQuizItems();

  const tabs = [
    { id: "all", label: "Semua" },
    { id: "assessment", label: "Assessment" },
    { id: "exam", label: "Ujian" },
  ];

  return (
    <View className="flex-1 bg-blue-100">
      <View className="flex-row justify-between items-center px-5 pt-14 pb-5 bg-white border-gray-200 rounded-b-3xl">
        <Text className="text-2xl font-bold text-slate-900">
          Quiz & Evaluasi
        </Text>
      </View>

      <ScrollView className="flex-1 p-5">
        <View className="bg-white px-5 py-3 border-b border-gray-200 rounded-3xl mb-5">
          <ScrollView horizontal showsHorizontalScrollIndicator={false}>
            <View className="flex-row gap-2">
              {tabs.map((tab) => (
                <TouchableOpacity
                  key={tab.id}
                  className={`px-4 py-2 rounded-2xl ${
                    selectedTab === tab.id ? "bg-blue-600" : "bg-gray-100"
                  }`}
                  onPress={() => setSelectedTab(tab.id)}
                >
                  <Text
                    className={`text-sm font-semibold ${
                      selectedTab === tab.id ? "text-white" : "text-gray-600"
                    }`}
                  >
                    {tab.label}
                  </Text>
                </TouchableOpacity>
              ))}
            </View>
          </ScrollView>
        </View>

        {loading ? (
          <View className="flex-1 justify-center items-center py-10">
            <ActivityIndicator size="large" color="#3b82f6" />
            <Text className="mt-3 text-sm text-gray-600">Memuat data...</Text>
          </View>
        ) : filteredItems.length === 0 ? (
          <View className="bg-white p-6 rounded-2xl items-center">
            <FileText size={32} color="#9ca3af" />
            <Text className="text-gray-600 text-sm mt-3">
              Belum ada{" "}
              {selectedTab === "exam"
                ? "ujian"
                : selectedTab === "assessment"
                  ? "assessment"
                  : "quiz"}
            </Text>
          </View>
        ) : (
          filteredItems.map((item) => {
            const isExam = item.type === "exam";
            let deadlineDate: Date;
            let isExpired = false;
            let daysLeft = 0;

            if (isExam) {
              const exam = item as ExamWithClass;
              deadlineDate = new Date(exam.start_at);
              const now = new Date();
              isExpired = exam.status === "published" && deadlineDate < now;
              daysLeft = Math.ceil(
                (deadlineDate.getTime() - now.getTime()) / (1000 * 60 * 60 * 24)
              );
            } else {
              const assessment = item as AssessmentWithClass;
              deadlineDate = new Date(assessment.deadline_at);
              const now = new Date();
              isExpired = deadlineDate < now;
              daysLeft = Math.ceil(
                (deadlineDate.getTime() - now.getTime()) / (1000 * 60 * 60 * 24)
              );
            }

            return (
              <TouchableOpacity
                key={`${item.type}-${isExam ? (item as ExamWithClass).id : (item as AssessmentWithClass).id}`}
                className="bg-white rounded-2xl p-4 mb-4"
              >
                <View className="flex-row justify-between items-center mb-3">
                  <View
                    className={`px-3 py-1.5 rounded-xl ${
                      isExam ? "bg-green-100" : "bg-blue-100"
                    }`}
                  >
                    <View className="flex-row items-center gap-1">
                      {isExam && <Award size={12} color="#10b981" />}
                      <Text
                        className={`text-xs font-semibold ${
                          isExam ? "text-green-600" : "text-blue-600"
                        }`}
                      >
                        {isExam ? "Ujian" : "Assessment"}
                      </Text>
                    </View>
                  </View>
                  <View className="flex-row items-center gap-1">
                    {isExpired ? (
                      <CheckCircle2 size={14} color="#10b981" />
                    ) : (
                      <AlertCircle size={14} color="#f59e0b" />
                    )}
                    <Text
                      className={`text-xs font-semibold ${
                        isExpired ? "text-green-500" : "text-amber-500"
                      }`}
                    >
                      {isExpired ? "Selesai" : `${daysLeft} hari lagi`}
                    </Text>
                  </View>
                </View>

                <Text className="text-base font-bold text-slate-900 mb-1">
                  {isExam
                    ? (item as ExamWithClass).name
                    : (item as AssessmentWithClass).name}
                </Text>
                <Text className="text-sm text-gray-600 mb-3">
                  {item.className}
                </Text>

                <View className="flex-row gap-4 mb-3">
                  <View className="flex-row items-center gap-1.5">
                    <Clock size={16} color="#64748b" />
                    <Text className="text-xs text-gray-600">
                      {deadlineDate.toLocaleDateString("id-ID", {
                        month: "short",
                        day: "numeric",
                      })}
                    </Text>
                  </View>
                  {isExam && (
                    <View className="flex-row items-center gap-1.5">
                      <Award size={16} color="#64748b" />
                      <Text className="text-xs text-gray-600">
                        {(item as ExamWithClass).duration} menit
                      </Text>
                    </View>
                  )}
                </View>

                <View className="flex-row gap-2">
                  <TouchableOpacity
                    className={`flex-1 py-2.5 rounded-2xl items-center ${
                      isExam ? "bg-green-600" : "bg-blue-600"
                    }`}
                  >
                    <Text className="text-white text-sm font-semibold">
                      Kerjakan
                    </Text>
                  </TouchableOpacity>
                  <TouchableOpacity className="px-5 py-2.5 rounded-2xl border border-gray-200 items-center">
                    <Text className="text-gray-600 text-sm font-semibold">
                      Detail
                    </Text>
                  </TouchableOpacity>
                </View>
              </TouchableOpacity>
            );
          })
        )}
      </ScrollView>

      {!isStudent && (
        <TouchableOpacity className="absolute bottom-5 right-5 w-14 h-14 rounded-full bg-blue-600 justify-center items-center shadow-lg">
          <Plus size={24} color="#ffffff" />
        </TouchableOpacity>
      )}
    </View>
  );
}
