import { format, parseISO } from "date-fns";
import { ja } from "date-fns/locale";

export const formatToJST = (dateString: string): string => {
    if (!dateString) return "";

    const date = parseISO(dateString);

    return format(date, "yyyy/MM/dd", { locale: ja });
};
